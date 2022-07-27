package main

import (
        "fmt"
        "time"
)

func main() {
        ret := 0
        start := time.Now()
        for i := 0; i <= 25; i++ {
                    ret = fib(i)
                    fmt.Printf("fib(%d) = %d\n", i, ret)
        }

        end := time.Now()
        delta := end.Sub(start)
        fmt.Printf("Time consume is : %s\n", delta)
}

func fib(n int) (ret int) {
        if n <= 1 {
                ret = 1
        } else {
                ret = fib(n-1) + fib(n-2)
        }
        return
}
