package main

import (
    "fmt"
)

func main() {
    test2()
}

func test1() {
    var s fmt.Stringer
    //s = "s"
    fmt.Println(s)
}

func test2() {
    const X = 7.0
    var x interface{} = X

    if y, ok := x.(int); ok {
        fmt.Println(y)
    } else {
        fmt.Println(int(y))
    }
}
