package main

import "fmt"

func main() {
        // defer存在的意义在于释放某些已分配的资源
        fmt.Println("-------withDefer--------\n")
        first()
        fmt.Println("-------withNoDefer--------\n")
        withNoDefer()
}

func withNoDefer() {
        fmt.Println("I am the first.\n")
        second()
        fmt.Println("I am the third.\n")
}

func first() {
        fmt.Println("I am the first.\n")
        defer second()
        fmt.Println("I am the third.\n")
}

func second() {
        fmt.Println("I am the second.\n")
}
