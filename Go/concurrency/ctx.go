package main

import (
    "fmt"
    "context"
    "time"
)

func monitor(ctx context.Context, num int) {
    for {
        select {
        case v := <- ctx.Done():
            fmt.Printf("Monitor %v recv %v.\n", num, v)
            return
        default:
            fmt.Printf("Monitor %v monitoring...\n", num)
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    for i := 1; i <= 5; i++ {
        go monitor(ctx, i)
    }

    time.Sleep(time.Second)
    cancel()

    time.Sleep(5 * time.Second)
    fmt.Println("Main program exited!")
}
