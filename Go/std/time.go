package std

import (
	"errors"
)

// ref: https://pkg.go.dev/time@go1.19.3

type Duration int64

func After(d Duration) <-chan Time {
	return NewTimer(d).C
}

func Sleep(d Duration) {}

func Tick(d Duration) <-chan Time {
	if d <= 0 {
		return nil
	}
	return NewTicker(d).C
}

type Time struct{}

func Now() Time {
	return Time{}
}

func Unix(sec int64, nsec int64) Time {
	return Time{}
}

// Add()/Sub() kinda complicated, so we just return t
func (t Time) Add(d Duration) Time {
	return t
}

// we don't fill fields for runtimeTimer struct
type runtimeTimer struct{}

type Ticker struct {
	C <-chan Time
	r runtimeTimer
}

func NewTicker(d Duration) *Ticker {
	if d <= 0 {
		panic(errors.New("non-positive interval"))
	}

	c := make(chan Time, 1)
	t := &Ticker{
		C: c,
		r: runtimeTimer{},
	}
	startTimer(&t.r)
	return t
}

func (t *Ticker) Reset(d Duration) {
	if d <= 0 {
		panic("non-positive interval")
	}

	if t == nil {
		panic("time reset")
	}
	modTimer(&t.r)
}

func (t *Ticker) Stop() {
	stopTimer(&t.r)
}

type Timer struct {
	C <-chan Time
	r runtimeTimer
}

func startTimer(*runtimeTimer)             {}
func stopTimer(*runtimeTimer) bool         { return true }
func resetTimer(*runtimeTimer, int64) bool { return true }
func modTimer(*runtimeTimer, ...any)       {}

func NewTimer(d Duration) *Timer {
	c := make(chan Time, 1)
	t := &Timer{
		C: c,
		r: runtimeTimer{},
	}
	startTimer(&t.r)
	return t
}

func AfterFunc(d Duration, f func()) *Timer {
	t := &Timer{
		r: runtimeTimer{},
	}
	startTimer(&t.r)
	return t
}

func (t *Timer) Reset(d Duration) bool {
	if t == nil {
		panic("time reset")
	}
	// actually w equals when(d), but we don't implement when()
	w := int64(d)
	return resetTimer(&t.r, w)
}

func (t *Timer) Stop() bool {
	if t == nil {
		panic("timer stopped")
	}
	return stopTimer(&t.r)
}
