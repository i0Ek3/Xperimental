package main

import (
	"fmt"
	"time"
)

func main() {
	req := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		req <- i
	}
	close(req)

	limiter := time.Tick(200 * time.Millisecond)

	for r := range req {
		<-limiter
		fmt.Println("req:", r, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyReqests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyReqests <- i
	}
	close(burstyReqests)
	for r := range burstyReqests {
		<-burstyLimiter
		fmt.Println("req:", r, time.Now())
	}
}
