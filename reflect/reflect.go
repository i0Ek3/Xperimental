package main

import (
	"reflect"
)

func valueof(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func typeof(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
