package main

import (
	"context"
	"fmt"
	"sync"
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

func test() {
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

// ++++++++++++++++++++++++++++++++++++++

var wg sync.WaitGroup
var ch = make(chan struct{})

func method() {
	wg.Add(2)
	go printLetter()
	go printNumber()
	wg.Wait()
}

func printNumber() {
	defer wg.Done()
	for i := 1; i <= 26; i += 2 {
		<-ch
		fmt.Printf("%d%d", i, i+1)
		ch <- struct{}{}
	}
}

func printLetter() {
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i += 2 {
		fmt.Print(string(i), string(i+1))
		ch <- struct{}{}
		<-ch
	}
}

func method2() {
	number := make(chan struct{})
	letter := make(chan struct{})

	go func() {
		i := 1
		for {
			// pause, wait number generate
			<-number
			// generate numbers
			fmt.Printf("%d%d", i, i+1)
			i += 2
			// notify letter goroutine to print
			letter <- struct{}{}
		}
	}()

	// use WaitGroup to wait goroutines finished
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 'A'; i <= 'Z'; i += 2 {
			// pause, wait letter generate
			<-letter
			fmt.Print(string(i), string(i+1))
			// notify number goroutine to print
			number <- struct{}{}
		}
	}(&wg)

	number <- struct{}{}
	// wait goroutines to exit
	wg.Wait()
}

func main() {
	// AB12
	method()
	fmt.Println("")
	// 12AB
	method2()
}
