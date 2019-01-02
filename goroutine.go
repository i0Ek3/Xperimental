package main

import (
        "fmt"
        "time"
)

func main() {
        fmt.Println("Running in main()...")
        go long()
        go short()
        fmt.Println("Sleeping in main()...")
        time.Sleep(10 * 1e9)
        fmt.Println("Ending in main()!")
}

func long() {
        fmt.Println("Running in long()...")
        time.Sleep(5 * 1e9)
        fmt.Println("Ending in long()!")
}

func short() {
        fmt.Println("Running in short()...")
        time.Sleep(2 * 1e9)
        fmt.Println("Ending in short()!")
}
