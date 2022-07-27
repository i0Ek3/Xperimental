package main 

import (
        "fmt"
        "errors"
)

var errNotFound error = errors.New("Not Found!")

func main() {
        fmt.Printf("error: %v", errNotFound)
}

