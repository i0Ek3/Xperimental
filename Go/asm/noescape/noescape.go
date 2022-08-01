package noescape

//go:noescape
func noescape(d []byte) (b []byte)

func test() int {
	var buf [1024]byte
	data := noescape(buf[:])
	return len(data)
}
