package main

import (
        "fmt"
        "time"
)

func test() {
        fmt.Println("Can you show me?")
}

func main() {
        go test()
        fmt.Println("hi, this is a test.")
        time.Sleep(time.Second)
}
