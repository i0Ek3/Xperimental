package main

import (
        "fmt"
        "time"
)

func main() {
        fmt.Println("-----------use channel--------")
        ch := make(chan string)
        go sendData(ch)
        go getData(ch)
        time.Sleep(1e9)

        fmt.Println("\n---------use goruntines-------")
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

func sendData(ch chan string) {
        ch <- "A"
        ch <- "B"
        ch <- "C"
        ch <- "D"
        ch <- "E"
}

func getData(ch chan string) {
        //var input string
        for {
                input := <- ch
                fmt.Printf("%s ", input)
        }
}


