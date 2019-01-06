package main

import (
        "fmt"
        "time"
)

// select 实现了一种监听模式，通常用在无限循环中；在某些情况下，可以用过 break 语句退出循环。

func main() {
        ch1 := make(chan int)
        ch2 := make(chan int)
        ch3 := make(chan int)
        
        go pump1(ch1)
        go pump2(ch2)
        go pump3(ch3)
        go suck(ch1, ch2, ch3)

        time.Sleep(1e9)
}

func pump1(ch chan int) {
        for i := 0; ; i++ {
                ch <- i * 2
        }
}

func pump2(ch chan int) {
        for i := 0; ; i++ {
                ch <- i + 3
        }
}

func pump3(ch chan int) {
        for i := 0; ; i++ {
                ch <- i * 0
        }
}


func suck(ch1, ch2, ch3 chan int) {
        for {
                select {
                case v := <- ch1:
                        fmt.Printf("Channel 1: received %d\n", v)
                case v := <- ch2:
                        fmt.Printf("Channel 2: received %d\n", v)
                case v := <- ch3:
                        fmt.Printf("Channel 3: received %d\n", v)
                }
        }
}
