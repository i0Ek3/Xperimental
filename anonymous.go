package main 

import (
        "fmt"
)

func main() {
    f()
    hyber()
}

func f() {
        noname := func(a int, b int, c int) int { return a+b*c }(3, 5, 0)
        fmt.Printf("noname = %d, it's address is: %p\n", noname, &noname)
}

func hyber() {
        defer func() {
                fmt.Printf("Print me first.\n")
        }()
        fmt.Println("Print me second.\n")
}



