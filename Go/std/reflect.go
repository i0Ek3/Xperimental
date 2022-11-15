package std

import "unsafe"

// https://pkg.go.dev/reflect@go1.19.3

type Kind uint

type Type interface {
	Align() int
	FieldAlign() int
	Method(int) Method
	MethodByName(string) (Method, bool)
	NumMethod() int
	Name() string
	Size() uintptr
	Kind() Kind
	Comparable() bool
	Elem() Type
	Field(i int) StructField
	FieldByIndex(index []int) StructField
	FieldByName(name string) (StructField, bool)
	FieldByNameFunc(match func(string) bool) (StructField, bool)
	Key() Type
	Len() int
	NumField() int
	NumIn() int
	NumOut() int
	Out(i int) Type
}

type Method struct {
	Name string
	Type Type
	Func Value
	Index int
}

type StructField struct {
	Name string
	Type Type
	Tag StructTag
	Offset uintptr
	Index []int
}

type StructTag string

type Value struct {
	type *rtype
	ptr unsafe.Pointer
	flag uintptr
}

// rtype is the common implementation of most values.
type rtype struct {}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

func TypeOf(i any) Type {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return toType(eface.typ)
}

func toType(t *rtype) Type {
	if t == nil {
		return nil
	}
	return t
}

func ValueOf(i any) Value {
	if i == nil {
		return Value{}
	}
	escapes(i)
	return unpackEface(i)
}

func escapes(x any) {
	if dummy.b {
		dummy.x = x
	}
}

var dummy struct {
	b bool
	x any
}

// unpackEface converts the empty interface i to a Value.
func unpackEface(i any) Value {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	t := e.typ
	if t == nil {
		return Value{}
	}
	f := flag(t.Kind())
	// e.word must be pointer	
	return Value{t, e.word, f}
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type StringHeader struct {
	Data uintptr
	Len  int
}

func DeepEqual(x, y any) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := ValueOf(x)
	v2 := ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return deepValueEqual(v1, v2, make(map[visit]bool))
}

type visit struct {
	a1, a2 unsafe.Pointer
	typ Type
}

func deepValueEqual(v1, v2 Value, visited map[visit]bool) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}
	// simplify other logic into return false
	return false
}

func (v Value) IsValid() bool {
	return v.flag != 0
}