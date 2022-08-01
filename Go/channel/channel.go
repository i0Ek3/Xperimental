package main

import (
	"fmt"
	"time"
)

// do not close uninitialized channel, it will panic
func test() {
	var ch chan int
	close(ch)
}

// do not close one channel twice, it will panic
func test1() {
	ch := make(chan int, 1)
	ch <- 1
	num, ok := <-ch
	if ok {
		close(ch)
	}
	fmt.Println(num)
	close(ch)
}

type (
	Receiver <-chan int
	Sender   chan<- int
)

func test2() {
	ch := make(chan int)

	go func() {
		var send Sender = ch
		fmt.Println("Sending... 100")
		send <- 100
	}()

	go func() {
		var recv Receiver = ch
		value := <-recv
		fmt.Printf("Receiving... %d", value)
	}()
	time.Sleep(time.Second)
	close(ch)
}

func test3() {
	ch := make(chan int, 10)

	go func(c chan int) {
		n := cap(c)
		x, y := 1, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// do not send msg to closed channel, it will panic
func test4() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	ch <- 2
}

// read msg from closed channel cannot panic
// if there is msg, it will read it out
// if not, it will read zeor value of this channel
func test5() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	fmt.Println(<-ch)
	v, ok := <-ch
	fmt.Println(v, ok)
}

func test6() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	go func() {
		for i := 0; i < 10; i++ {
			v, ok := <-ch
			fmt.Println(v, ok)
		}
	}()
	time.Sleep(time.Second)
}

func test7() {
	msg := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			msg <- i
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(<-msg)
	}
}

func test8() {
	w := func(done chan bool) {
		fmt.Print("doing...")
		time.Sleep(time.Second)
		fmt.Println("done!")
		done <- true
	}
	done := make(chan bool, 1)
	go w(done)
	<-done
}

func test9() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			//default:
		}
	}
}

func test10() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "res 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "res 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func test11() {
	msg := make(chan string)
	sig := make(chan bool)

	select {
	case m := <-msg:
		fmt.Println(m)
	default:
		fmt.Println("nothing")
	}

	str := "hi"
	select {
	case msg <- str:
		fmt.Println("sent msg", msg)
	default:
		fmt.Println("nothing")
	}

	select {
	case m := <-msg:
		fmt.Println("received msg", m)
	case s := <-sig:
		fmt.Println("received sig", s)
	default:
		fmt.Println("nothing")
	}
}

func test12() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, ok := <-jobs
			if ok {
				fmt.Println("received job", j)
			} else {
				fmt.Println("channel was closed")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

func test13() {
	ch := make(chan int)

	pump := func(ch chan int) {
		for i := 0; ; i++ { // dead loop
			ch <- i
		}
	}
	go pump(ch)

	suck := func(ch chan int) {
		for {
			fmt.Println(<-ch)
		}
	}
	go suck(ch)

	time.Sleep(1e9)
	//fmt.Println(<-ch) // no reciver

	fmt.Println("\n----------------------------")
	// panic: all goruntines asleep, deadloop

	out := make(chan int)
	out <- 2

	f := func(in chan int) {
		fmt.Println(<-in)
	}
	go f(out)
}

func test14() {
	plusone := func(ch chan bool, x *int) {
		ch <- true
		*x = *x + 1
		<-ch
	}
	var x int
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		go plusone(ch, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("x = ", x)
}

func test15() {
    c1 := make(chan int, 1000)
    
    // c1 is empty, deadlock
    /*for v := range c1 {
        fmt.Println(v)
    }
    //close(c1)
    */
    
    // main is so fast in order that go cannot excute
    // so this part cannot deadlock
    go func() {
        // c1 is empty
        for v := range c1 {
            fmt.Println(v)
        }
    }()
}

func main() {
	test15()
}
