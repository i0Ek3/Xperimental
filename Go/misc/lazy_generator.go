package main

import (
        "fmt"
)

// 惰性生成器
// 生成器是指当被调用时返回一个序列中下一个值的函数，这种特性也称为惰性求值

var resume chan int

func integers() chan int {
        yield := make(chan int)
        count := 0
        go func() {
                for {
                        yield <- count
                        count++
                }
        }()
        return yield 
}

func generateInteger() int {
        return <- resume 
}

func main() {
        resume = integers()
        fmt.Println(generateInteger())
        fmt.Println(generateInteger())
        fmt.Println(generateInteger())
}




