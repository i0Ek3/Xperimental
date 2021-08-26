package main

import (
        "fmt"
        "sort"
)

var (
        seq = map[string]int{
                "a": 1,
                "c": 3,
                "b": 2,
                "d": 4,
                "f": 6,
                "e": 5}
)

func sorted() {
        fmt.Println("Original Seq:")
        for k, v := range seq {
                fmt.Printf("Key = %v, Val = %v\n", k, v)
        }
        keys := make([]string, len(seq))
        i := 0
        for k, _ := range seq {
                keys[i] = k
                i++
        }
        sort.Strings(keys)
        fmt.Println()
        fmt.Println("Sorted Seq:")
        for _, k := range keys {
                fmt.Printf("Key = %v, Val = %v\n", k, seq[k])
        }
}

func inverted() {
        invertMap := make(map[int]string, len(seq)) 
        for k, v := range seq {
                invertMap[v] = k
        }
        fmt.Println("Inverted:")
        for k, v := range invertMap {
                fmt.Printf("Key = %v, Val = %v\n", k, v)
        }

}

func main() {
        sorted()
        inverted()
}

