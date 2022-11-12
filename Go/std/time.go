package main

import (
	"errors"
	"fmt"
	"time"
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

func statusUpdate() string { return "" }
func TestTick() {
	t := time.Tick(3 * time.Second)
	for next := range t {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

func handle(int) {}
func TestAfterSleep() {
	var c chan int
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Second):
		time.Sleep(100 * time.Microsecond)
		fmt.Println("timed out")
	}
}

func TestParseDuration() {
	h, _ := time.ParseDuration("10h")
	c, _ := time.ParseDuration("1h10m10s")
	m, _ := time.ParseDuration("1µs")
	n, _ := time.ParseDuration("1us")
	fmt.Println(h, c, m, n)
}

func TestNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println("current time:", t)
		}
	}
}

func TestAdd() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenS := start.Add(10 * time.Second)
	afterTenH := start.Add(10 * time.Hour)
	fmt.Printf("start = %v\n", start)
	fmt.Printf("after 10s = %v\n", afterTenS)
	fmt.Printf("after 10h = %v\n", afterTenH)
}

func TestUnix() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)
}

func main() {
	recap()
}

func recap() {
	o := fmt.Println
	o(time.Now())
	o(time.Now().Unix())
	o(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	o(time.ParseDuration("1s2m3h"))
}
