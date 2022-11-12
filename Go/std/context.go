package main

// ref: https://cs.opensource.google/go/go/+/go1.19.3:src/context/context.go

import (
	"context"
	"errors"
	"fmt"
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

func TestWithCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

const shortDur = 1 * time.Millisecond

func TestWithDeadline() {
	d := time.Now().Add(shortDur)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDur)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithValue() {
	type CtxKey string
	f := func(ctx context.Context, k CtxKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found:", v)
			return
		}
		fmt.Println("not found:", k)
	}

	k := CtxKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, CtxKey("color"))
}

func main() {
	//TestWithValue()
	recap()
}

func recap() {
	key := "hi"
	d := time.Now().Add(shortDur)
	ctx := context.WithValue(context.Background(), key, "there")
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, shortDur)
	ctx, cancel = context.WithDeadline(ctx, d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("reason:", ctx.Err())
	}
}
