package main 

import (
        "fmt"
)

func main() {
        cal(10, div)
}

func div(a, b int) {
        fmt.Printf("%d divided %d is: ", a, b, a/b)
}

func sub(a, b int) {
        fmt.Printf("%d sub %d is: ", a, b, a-b)
}

func cal(m int, f func(int, int)) {
        f(m, 1)
}
