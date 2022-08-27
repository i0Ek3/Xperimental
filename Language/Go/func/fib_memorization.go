package main

import (
        "fmt"
        "time"
)

const LIM = 25

var fibs [LIM]uint64

func main() {
        var ret uint64 = 0
        start := time.Now()
        for i := 0; i < LIM; i++ {
                ret = fib(i)
                fmt.Printf("fib(%d) = %d\n", i, ret)
        }
        end := time.Now()
        delta := end.Sub(start)
        fmt.Printf("Time consume is: %s\n", delta)
}

func fib(n int) (ret uint64) {
        if fibs[n] != 0 {
                ret = fibs[n]
                return 
        } 
        if n <= 1 {
                ret = 1
        } else {
                ret = fib(n-1) + fib(n-2)
        }
        fibs[n] = ret
        return 
}

