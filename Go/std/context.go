package std

// ref: https://cs.opensource.google/go/go/+/go1.19.3:src/context/context.go

import (
	"errors"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key any) any {
	return nil
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

type CancelFunc func()

var Canceled = errors.New("canceled")

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	if parent == nil {
		panic("nil parent")
	}

	c := newCancelCtx(parent)
	//propagateCancel(parent, &c)

	return &c, nil
}

type cancelCtx struct {
	Context
}

func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

type timerCtx struct {
	cancelCtx
	deadline time.Time
}

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if parent == nil {
		panic("nil parent")
	}

	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}

	return c, nil
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("nil parent")
	}

	if key == nil {
		panic("nil key")
	}
	/*if !reflectlite.TypeOf(key).Compareable() {
		panic("key must be compareable")
	}*/

	return &valueCtx{parent, key, val}
}

type valueCtx struct {
	Context
	key, val any
}
