package main

import (
    "fmt"
    "strings"
)

func main() {
    var oldString string = "abc "
    var newString string = "deffmuf"
    var new string;

    new = strings.Repeat(oldString, strings.Count(newString, "f"))
    fmt.Printf("The new string repeated is: %s, and the first letter f's index is %d.\n ", new, strings.Index(newString, "f"))
    fmt.Printf("And the upper letter of oldString is %s.\n", strings.ToUpper(oldString))
    fmt.Printf("And after trim the blank of oldString is %s.\n", strings.TrimSpace(oldString))


}
