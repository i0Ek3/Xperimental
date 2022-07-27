package main

import(
    "fmt"
    "time"
)

func plusone(ch chan bool, x *int) {
    ch <- true
    *x = *x + 1
    <- ch
}

func main() {
    var x int
    ch := make(chan bool, 1)
    for i := 0; i < 1000; i++ {
        go plusone(ch, &x)
    }
    time.Sleep(time.Second)
    fmt.Println("x = ", x)
}
