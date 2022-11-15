package std

// https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/unsafe/unsafe.go

type ArbitraryType int
type Pointer *ArbitraryType

func Sizeof(x ArbitraryType) uintptr  { return uintptr(x) }
func Alignof(x ArbitraryType) uintptr { return uintptr(x) }
