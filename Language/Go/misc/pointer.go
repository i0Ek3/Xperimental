package main

import "fmt"

func intPointer() {
    var int1 = 10
    fmt.Printf("An integer %d its location in memory is %p.\n", int1, &int1)
    var intP *int
    intP = &int1
    fmt.Printf("The value at memory location %p is %d.\n", intP, *intP)

}

func stringPointer() {
    s := "Merci~"
    fmt.Printf("Here is the original string s: %s.\n", s)
    var pstr *string = &s
    *pstr = "Garacias~"
    fmt.Printf("Here is the pointer pstr: %p.\n", pstr)
    fmt.Printf("Here is the string *pstr: %s.\n", *pstr)
    fmt.Printf("Here is the string s: %s.\n", s)

}

func main() {
    intPointer()
    stringPointer()
}
