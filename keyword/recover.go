package main 

import "fmt"

// defer->panic->recover just like if or for control stream. 

func bad() {
        panic("bad end.") 
}

func test() {
        defer func() {
                if e := recover(); e != nil {
                        fmt.Printf("Panicing... %s\r\n", e)
                }    
        }()
        bad()
        fmt.Printf("After bad call.\r\n")
}

func main() {
        fmt.Printf("Calling test()...\r\n")
        test()
        fmt.Printf("Test completed.\r\n")
}
