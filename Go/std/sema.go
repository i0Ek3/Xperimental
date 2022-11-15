package std

import (
	"container/list"
	"context"
	"sync"
)

// sema provides a weighted semaphore implementation.

type waiter struct {
	n     int64
	ready chan<- struct{}
}

type Weighted struct {
	size    int64
	cur     int64
	mu      sync.Mutex
	waiters list.List
}

func NewWeighted(n int64) *Weighted {
	w := &Weighted{size: n}
	return w
}

func (w *Weighted) Acquire(ctx context.Context, n int64) error {
	w.mu.Lock()
	if w.size-w.cur >= n && w.waiters.Len() == 0 {
		w.cur += n
		w.mu.Unlock()
		return nil
	}
	if n > w.size {
		w.mu.Unlock()
		<-ctx.Done()
		return ctx.Err()
	}

	ready := make(chan struct{})
	wait := waiter{n: n, ready: ready}
	elem := w.waiters.PushBack(wait)
	w.mu.Unlock()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		w.mu.Lock()
		select {
		case <-ready:
			err = nil
		default:
			isFront := w.waiters.Front() == elem
			w.waiters.Remove(elem)
			// 目前存在可用的 token
			if isFront && w.size > w.cur {
				w.notifyWaiters()
			}
		}
		w.mu.Unlock()
		return err
	case <-ready:
		return nil
	}
}

// like TryLock()
func (w *Weighted) TryAcquire(n int64) bool {
	w.mu.Lock()
	success := w.size-w.cur >= n && w.waiters.Len() == 0
	if success {
		w.cur += n
	}
	w.mu.Unlock()
	return success
}

func (w *Weighted) Release(n int64) {
	w.mu.Lock()
	w.cur -= n
	if w.cur < 0 {
		w.mu.Unlock()
		panic("semaphore: not enough token can be released")
	}
	w.notifyWaiters()
	w.mu.Unlock()
}

func (w *Weighted) notifyWaiters() {
	for {
		next := w.waiters.Front()
		// no more waiters
		if next == nil {
			break
		}
		wait := next.Value.(waiter)
		// no enough tokens
		if w.size-w.cur < wait.n {
			break
		}
		w.cur += wait.n
		w.waiters.Remove(next)
		close(wait.ready)
	}
}
