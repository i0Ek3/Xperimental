package main

import (
    "fmt"
    "time"
)

const n = 10

var (
    num = make(chan struct{}, 1)
    alp = make(chan struct{}, 1)
)

func printNum() {
    for i := 1; i <= n; i++ {
        <-alp
        fmt.Println(i)
        num <- struct{}{}
    }
}

func printAlp() {
    for i := 'A'; i < 'Z'; i++ {
        <-num
        fmt.Println(string(i))
        alp <- struct{}{}
    }
}

func main() {
    num <- struct{}{}
    go printNum()
    go printAlp()
    time.Sleep(time.Second)
}

