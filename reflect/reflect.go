package main

import (
    "fmt"
	"reflect"
)

func valueof(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func typeof(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

type Test struct {
    name string
    desc string
}

func main() {
    t := Test{}
    typ := reflect.TypeOf(t)
    fmt.Println(t)
    fmt.Println(typ.Kind())
}
