package std

// https://pkg.go.dev/sort@go1.19.3

func Slice(x any, less func(i, j int) bool) {}

type Interface interface {
	Len() IntSlice
	Less(i, j int) bool
	Swap(i, j int)
}

type F64Slice []float64
type IntSlice []int
type StringSlice []string

type lessSwap struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}
