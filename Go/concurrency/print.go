package main

import (
	"context"
	"fmt"
	"time"
)

func test1(in chan struct{}) {
	time.Sleep(time.Second)
	in <- struct{}{}
}

func test2(in chan struct{}) {
	time.Sleep(3 * time.Second)
	in <- struct{}{}
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		go test1(ch1)
		select {
		case <-ctx.Done():
			fmt.Println("test1 timeout")
			break
		case <-ch1:
			fmt.Println("test1 done")
		}
	}()

	go func() {
		go test2(ch2)
		select {
		case <-ctx.Done():
			fmt.Println("test2 timeout")
			break
		case <-ch2:
			fmt.Println("test2 done")
		}
	}()
	time.Sleep(5 * time.Second)
}
