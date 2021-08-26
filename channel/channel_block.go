package main

import (
        "fmt"
        "time"
)

func main() {
        ch := make(chan int)
        go pump(ch)
        go suck(ch)
        time.Sleep(1e9)
        //fmt.Println(<-ch) // no reciver

        fmt.Println("\n----------------------------")
        // panic: all goruntines asleep, deadloop

        out := make(chan int)
        out <- 2
        go f(out)
}

func pump(ch chan int) {
        for i := 0; ; i++ { // dead loop
                ch <- i
        }
}

func suck(ch chan int) {
        for {
                fmt.Println(<-ch)
        }
}

func f(in chan int) {
        fmt.Println(<-in)
}
