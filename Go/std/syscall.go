package std

// https://pkg.go.dev/syscall@go1.19.3

type Signal int

func (s Signal) Signal() {}

// use trap
func Kill(pid int, sig Signal) (err error) { return }
