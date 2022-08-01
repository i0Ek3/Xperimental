package main

import (
        "fmt"
        "time"
)

func test() {
        fmt.Println("Can you show me?")
}

func runtest() {
        go test()
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

func main() {
    //runtest()
    gotest()
}
