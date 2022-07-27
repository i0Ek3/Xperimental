package main

import(
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
    num, ok := <- ch; if ok {
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
    v, ok := <- ch
    fmt.Println(v, ok)
}

func test6() {
    ch := make(chan int, 10)
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)

    go func(){
        for i := 0; i < 10; i++ {
            v, ok := <- ch
            fmt.Println(v, ok)
        }
    }()
    time.Sleep(time.Second)
}

func main() {
    test6()
}
