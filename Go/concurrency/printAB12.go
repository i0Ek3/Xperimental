package main

import (
    "fmt"
    "sync"
)

func main() {
    number := make(chan bool)
    letter := make(chan bool)

    go func() {
        i := 1
        for {
            // pause, wait number generate
            <-number
            // generate numbers
            fmt.Printf("%d%d", i, i+1)
            i += 2
            // notify letter goroutine to print
            letter<-true
        }
    }()

    // use WaitGroup to wait goroutines finished
    wg := sync.WaitGroup{}
    wg.Add(1)

    go func(wg *sync.WaitGroup) {
        defer wg.Done()
        for i := 'A'; i <= 'Z'; i += 2 {
            // pause, wait letter generate
            <-letter
            fmt.Print(string(i), string(i+1))
            // notify number goroutine to print
            number<-true
        }
    }(&wg)

    // finished print action
    number<-true
    // wait goroutines to exit
    wg.Wait()
}
