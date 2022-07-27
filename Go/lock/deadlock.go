package main

import (
    "fmt"
)

func deadlock() {
    ch := make(chan string)
    fmt.Println(<-ch)
    ch <- "hello"
}

func hi(ch chan string) {
    <- ch
}

func test1() {
    ch := make(chan string)
    go hi(ch)
    ch <- "hi"
}

func test2() {
    ch := make(chan string)
    go func() {
        ch <- "hi"
        ch <- "Go"
        close(ch)
    }()

    for v := range ch {
        fmt.Println(v)
    }
}

func main() {
    test2()
}
