package main

import (
        "fmt"
        "errors"
        "math"
)

func main() {
        add, mul, sub := withName(12, 4)
        fmt.Println("Sum :", add, " Mulit :", mul, " Subtract :", sub)   
        add, mul, sub = noneName(12, 4)
        fmt.Println("Sum :", add, " Mulit :", mul, " Subtract :", sub)   

        fmt.Println("First example with -1: ")
        ret1, err1 := testSqrt(-1)
        if err1 != nil {
                fmt.Println("ERRRRRRRRRRRRR! Values are: ", ret1, err1)
        } else {
                fmt.Println("Ok! Values are: ", ret1, err1)
        } 
    
        fmt.Println("Second example with 3: ")
        if ret2, err2 := testSqrt2(3); err2 != nil {
                fmt.Println("ERRRRRRRRRRRRR! Values are: ", ret2, err2)
        } else {
                fmt.Println("Ok! Values are: ", ret2, err2)
        } 
        fmt.Println(testSqrt2(3))
    
}

func withName(a, b int) (int, int, int) {
        return a+b, a*b, a-b
}

func noneName(a, b int) (m int, n int, o int) {
        m, n, o = a+b, a*b, a-b
        return 
}

func testSqrt(f float64) (float64, error) {
        if f < 0 {
                return float64(math.NaN()), errors.New("Errors running, zero fucked.")
        }
        return math.Sqrt(f), nil
}

func testSqrt2(f float64) (ret float64, err error) {
        if f < 0 {
                ret = float64(math.NaN())
                err = errors.New("You get wrong! Here exist a negative value!")
        } else {
                ret = math.Sqrt(f)
        }
        return 
}

