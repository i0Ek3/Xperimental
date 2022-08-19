package main

import (
    "fmt"
    "sync"
)

func main() {
    number := make(chan struct{})
    letter := make(chan struct{})

    go func() {
        i := 1
        for {
            // pause, wait number generate
            <-number
            // generate numbers
            fmt.Printf("%d%d", i, i+1)
            i += 2
            // notify letter goroutine to print
            letter<-struct{}{}
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
            number<-struct{}{}
        }
    }(&wg)

    // finished print action
    number<-struct{}{}
    // wait goroutines to exit
    wg.Wait()
}
