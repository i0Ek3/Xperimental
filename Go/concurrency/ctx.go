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

func ctxTest() {
    ctx, cancel := context.WithCancel(context.Background())

    for i := 1; i <= 5; i++ {
        go monitor(ctx, i)
    }

    time.Sleep(time.Second)
    cancel()

    time.Sleep(5 * time.Second)
    fmt.Println("Main program exited!")
}

func ctxTest2() {
    ctx := context.Background()
    pass(ctx)

    ctx = context.WithValue(ctx, "passID", "i0Ek3")
    pass(ctx)
}

func pass(ctx context.Context) {
    passID, ok := ctx.Value("passID").(string)
    if ok {
        fmt.Printf("passID = %s\n", passID)
    } else {
        fmt.Println("passID is null")
    }
}

func main() {
    ctxTest2()
}
