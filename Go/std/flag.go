package std

// https://pkg.go.dev/flag@go1.19.3

import "io"

type Value interface {
	String() string
	Set(string) error
}

type Flag struct {
	Name     string
	Usage    string
	Value    Value
	DefValue string
}

func Bool(name string, value bool, usage string) *bool {
	// use given parameters to create a bool pointer
	// buf here we just use new(bool) to return
	return new(bool)
}

type ErrorHandling int

type FlagSet struct {
	Usage          func()
	name           string
	parsed         bool
	actual, formal map[string]*Flag
	args           []string
	errorHandling  ErrorHandling
	output         io.Writer
}
