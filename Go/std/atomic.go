package std

import "unsafe"

// https://pkg.go.dev/sync/atomic@go1.19.3

type noCopy struct{}
type align64 struct{}

type Bool struct {
	_ noCopy
	v uint32
}

// Load atomically loads and returns the value stored in x.
func (x *Bool) Load() bool { return LoadUint32(&x.v) != 0 }

func LoadUint32(addr *uint32) (val uint32) {
	return val
}

// Store atomically stores val into x.
func (x *Bool) Store(val bool) { StoreUint32(&x.v, b32(val)) }

func StoreUint32(addr *uint32, val uint32) {}

// Swap atomically stores new into x and returns the previous value.
func (x *Bool) Swap(new bool) (old bool) { return SwapUint32(&x.v, b32(new)) != 0 }

func SwapUint32(addr *uint32, new uint32) (old uint32) {
	return old
}

func (x *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return CompareAndSwapUint32(&x.v, b32(old), b32(new))
}

func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	return true
}

func b32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

type Value struct {
	v any
}

type Uintptr struct {
	_ noCopy
	v uintptr
}

type Int64 struct {
	_ noCopy
	_ align64
	v int64
}

type Uint64 struct {
	_ noCopy
	_ align64
	v uint64
}

type Pointer[T any] struct {
	_ noCopy
	v unsafe.Pointer
}
