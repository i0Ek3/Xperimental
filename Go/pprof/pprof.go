package main

import (
    "fmt"
    "runtime/pprof"
    "os"
)

func debug(filename string) {
    f, _ := os.OpenFile(filename, os.O_CREATE | os.O_RDWR, 0644)
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
}

func fib(n int) int {
    if n <= 1 {
        return 1
    }
    a, b := 1, 2
    for i := 2; i < n; i++ {
        sum := a + b
        a = b
        b = sum
    }
    return b
}

func main() {
    filename := "cpu.profile"
    debug(filename)

    n := 10
    for i := 1; i <= 5; i++ {
        fmt.Printf("fib(%d) = %d\n", n, fib(n))
        n += 3 * i
    }
}
