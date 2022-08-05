package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(out chan<- int) {
	for i := 0; ; i++ {
		out <- i * i
	}
}

func Consumer(input <-chan int) {
	for v := range input {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(ch)
	go Producer(ch)
	go Consumer(ch)

	//time.Sleep(time.Second * 3)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit %v\n", <-sig)

}
