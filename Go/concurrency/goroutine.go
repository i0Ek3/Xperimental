package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	test2()
	//gotest()
}

func test() {
	fmt.Println("Can you show me?")
}

func test1() {
	fmt.Println("-----------use channel--------")
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
	time.Sleep(1e9)

	fmt.Println("\n---------use goruntines-------")
	fmt.Println("Running in main()...")
	go long()
	go short()
	fmt.Println("Sleeping in main()...")
	time.Sleep(10 * 1e9)
	fmt.Println("Ending in main()!")

	go test()
	fmt.Println("hi, this is a test.")
}

func long() {
	fmt.Println("Running in long()...")
	time.Sleep(5 * 1e9)
	fmt.Println("Ending in long()!")
}

func short() {
	fmt.Println("Running in short()...")
	time.Sleep(2 * 1e9)
	fmt.Println("Ending in short()!")
}

func sendData(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"
	ch <- "E"
	close(ch)
}

func getData(ch chan string) {
	//var input string
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func canyou() {
	fmt.Println("Can you show me?")
}

func runtest() {
	go canyou()
	fmt.Println("hi, this is a test.")
	time.Sleep(time.Second)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func gotest() {
	f("direct")

	go func(msg string) {
		//fmt.Println(msg)
		f(msg)
	}("going")

	go f("goroutine")

	// if you comment this line, goroutine will not be exceuted
	time.Sleep(time.Second)
	fmt.Println("done")
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func test2() {
	var readOps, writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 100; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writesOps:", writeOpsFinal)
}
