package main

import (
        "fmt"
)

// close channel

func main() {
        ch := make(chan float64)
        
        //way 1
        defer close(ch)

        //way 2
        v, ok := <-ch
        if !ok {
                break
        }
        
        //way 3
        for input := range ch {
                //
        } 
}


