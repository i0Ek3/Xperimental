package main

import (
        "fmt"
        "reflect"
)

type TagType struct {
        field1 bool     "Whethter true or false."
        field2 string   "Things' name."
        field3 int      "Stuffs' price."
}

func main() {
        tt := TagType{true, "LV", 10}
        for i := 0; i < 3; i++ {
                refTag(tt, i)
        }
}

func refTag(tt TagType, ix int) {
        ttType := reflect.TypeOf(tt)
        ixField := ttType.Field(ix)
        fmt.Printf("%v\n", ixField.Tag)
}
