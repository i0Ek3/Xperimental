package std

// https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/runtime/symtab.go;drc=9839668b5619f45e293dd40339bf0ac614ea6bee;l=718

type Func struct {
	opaque struct{}
}

// FuncForPC kinda complicated, we just return the empty Func
func FuncForPC(pc uintptr) *Func {
	return &Func{}
}

type funcInfo struct {
	*_func
	datap *moduledata
}

// moduledata records information about the layout of the executable image.
// it's kinda complicated so we just use literal value of struct here
type moduledata struct{}

type Frame struct {
	PC       uintptr
	Func     *Func
	Function string
	File     string
	Line     int
	Entry    uintptr
	funcInfo funcInfo
}
