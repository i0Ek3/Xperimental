package std

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

// call 是在飞行过程中的或者已经完成的 Do()
type call struct {
	wg    sync.WaitGroup
	val   any
	err   error
	dups  int
	chans []chan<- Result
}

// Do() 的结果
type Result struct {
	Val    any
	Err    error
	Shared bool
}

// 表示一组调用
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

type panicError struct {
	value interface{}
	stack []byte
}

var errGoexit = errors.New("runtime.Goexit was called")

// 确保给定 key 仅执行一次
func (g *Group) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	// 有则取出返回
	if c, ok := g.m[key]; ok {
		c.dups++
		g.mu.Unlock()
		c.wg.Wait()

		if e, ok := c.err.(*panicError); ok {
			panic(e)
		} else if c.err == errGoexit {
			runtime.Goexit()
		}
		return c.val, c.err, true
	}
	// 无则新建
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// 调用具体的处理函数
	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}

// 同 Do 类似，不过用通道返回结果
func (g *Group) DoChan(key string, fn func() (any, error)) <-chan Result {
	ch := make(chan Result, 1)
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		c.chans = append(c.chans, ch)
		g.mu.Unlock()
		return ch
	}
	c := &call{chans: []chan<- Result{ch}}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	go g.doCall(c, key, fn)
	return ch
}

func (p *panicError) Error() string {
	return fmt.Sprintf("%v\n\n%s", p.value, p.stack)
}

func newPanicError(v any) error {
	st := debug.Stack()
	if line := bytes.IndexByte(st[:], '\n'); line >= 0 {
		st = st[line+1:]
	}
	return &panicError{value: v, stack: st}
}

func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
	normalReturn, recovered := false, false
	defer func() {
		if !normalReturn && !recovered {
			c.err = errGoexit
		}
		g.mu.Lock()
		defer g.mu.Unlock()
		c.wg.Done()
		// key 存在则删除
		if g.m[key] == c {
			delete(g.m, key)
		}
		if e, ok := c.err.(*panicError); ok {
			if len(c.chans) > 0 {
				go panic(e)
				select {}
			} else {
				panic(e)
			}
		} else if c.err == errGoexit {
		} else {
			for _, ch := range c.chans {
				ch <- Result{c.val, c.err, c.dups > 0}
			}
		}
	}()

	func() {
		defer func() {
			if !normalReturn {
				if r := recover(); r != nil {
					c.err = newPanicError(r)
				}
			}
		}()
		c.val, c.err = fn()
		normalReturn = true
	}()

	// 恢复捕获状态
	if !normalReturn {
		recovered = true
	}
}

// 直接安全地删除某个 key
func (g *Group) Forget(key string) {
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
}
