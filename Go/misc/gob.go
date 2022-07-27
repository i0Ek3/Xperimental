package main

import (
        "bytes"
        "fmt"
        "log"
        "encoding/gob" //use reflection in the progress of encode and decode
)

type P struct {
        X, Y, Z int
        Name string 
}

type Q struct {
        X, Y *int32
        Name string 
}

func main() {
        var network bytes.Buffer
        enc := gob.NewEncoder(&network)
        dec := gob.NewDecoder(&network)
        err := enc.Encode(P{3, 4, 5, "Kolin"})
        if err != nil {
            log.Fatal("Encode wrong: ", err)
        }
        var q Q
        err = dec.Decode(&q)
        if err != nil {
                log.Fatal("Decode wrong: ", err)
        }
        fmt.Printf("struct Q --> %q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

