package lock_chan

import "sync"

type ChanMutex struct {
	ch chan struct{}
}

func NewChanMutex() *ChanMutex {
	c := &ChanMutex{
		ch: make(chan struct{}, 1),
	}
	c.ch <- struct{}{}
	return c
}

func (c *ChanMutex) Lock() {
	<-c.ch
}

func (c *ChanMutex) Unlock() {
	c.ch <- struct{}{}
}

const TIMES = 100

func (c *ChanMutex) Add(v int) {
	for i := 0; i < TIMES; i++ {
		c.Lock()
		v += i
		c.Unlock()
	}
}

type CommonMutex struct {
	m sync.Mutex
}

func NewCommonMutex() *CommonMutex {
	return &CommonMutex{}
}

func (c *CommonMutex) Lock() {
	c.m.Lock()
}

func (c *CommonMutex) Unlock() {
	c.m.Unlock()
}

func (c *CommonMutex) Add(v int) {
	for i := 0; i < TIMES; i++ {
		c.Lock()
		v += i
		c.Unlock()
	}
}
