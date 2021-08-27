package main

import(
    "fmt"
    "time"
)

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

func main() {
    test3()
}
