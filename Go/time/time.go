package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}

func test1() {
	fmt.Println("wait 2 second")
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C
	fmt.Println("timer 1 fired")

	fmt.Println("wait 1 second")
	t2 := time.NewTimer(time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer 2 fired")
	}()

	fmt.Println("I'm just skip go func()")
	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("wait 2 second")
		fmt.Println("timer 2 stopped")
	}
	fmt.Println("2?")
	time.Sleep(2 * time.Second)
}

func test2() {
	tick := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool, 1)

	//done <- true
	go func() {
		fmt.Println("I'm the first.")
		for {
			select {
			case <-done:
				fmt.Println("Am I?")
				return
			case t := <-tick.C:
				fmt.Println("I'm the second.")
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("I'm the third.")
	tick.Stop()
	fmt.Println("I'm the fourth.")
	done <- true
	fmt.Println("tick stopped")
}
