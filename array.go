package main

import (
        "fmt"
)

func main() {
        var arr1 [5]int
        for i := 0; i < len(arr1); i++ {
                arr1[i] += i
        }

        for i := 0; i < len(arr1); i++ {
                fmt.Printf("Array at index %d is %d\n", i, arr1[i])
        }

        fmt.Printf("------------Up with for-----------down with for-range-------------\n")
        //for-range
        for i, _ := range arr1 {
                fmt.Printf("Array at index %d is %d\n", i, arr1[i])
        }
}
