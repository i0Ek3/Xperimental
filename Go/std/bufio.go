package std

import (
	"io"
)

// ref: https://pkg.go.dev/bufio@go1.19.3

const defaultBufSize = 4096
const minReadBufferSize = 16

type Reader struct {
	buf          []byte
	rd           io.Reader
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

func NewReaderSize(rd io.Reader, size int) *Reader {
	b, ok := rd.(*Reader)
	if ok && len(b.buf) >= size {
		return b
	}
	if size < minReadBufferSize {
		size = minReadBufferSize
	}
	r := new(Reader)
	r.reset(make([]byte, size), rd)
	return r
}

func (b *Reader) Reset(r io.Reader) {
	if b.buf == nil {
		b.buf = make([]byte, defaultBufSize)
	}
	b.reset(b.buf, r)
}

func (b *Reader) reset(buf []byte, r io.Reader) {
	*b = Reader{}
}

func NewReader(rd io.Reader) *Reader {
	return NewReaderSize(rd, defaultBufSize)
}

func (b *Reader) readErr() error {
	return nil
}

func (b *Reader) Buffered() int {
	return b.w - b.r
}

// 当缓存区有内容时，将缓存区内容全部填入p并清空缓存区
// 当缓存区没有内容的时候，且len(p)>len(buf)，即要读取的内容比缓存区还要大，直接去文件读取即可
// 当缓存区没有内容的时候，且len(p)<len(buf)，即要读取的内容比缓存区小，缓存区从文件读取内容充满缓存区，并将p填满（此时缓存区有剩余内容）
// 以后再次读取时缓存区有内容，将缓存区内容全部填入p并清空缓存区（此时和情况1一样）
func (b *Reader) Read(p []byte) (n int, err error) {
	n = len(p)

	// p 中没有数据
	if n == 0 {
		// 缓冲区中数量大于 0
		if b.Buffered() > 0 {
			return 0, b.readErr()
		}
		return 0, nil
	}

	//
	if b.r == b.w {
		// 2.1 读取过程中出错
		if b.err != nil {
			return 0, b.readErr()
		}
		// p 足够大
		if len(p) >= len(b.buf) {
			// 读取 p 中的数据
			n, b.err = b.rd.Read(p)
			if n < 0 {
				panic("non-positive")
			}
			// 读取到的数据长度大于 0，则更新索引
			if n > 0 {
				b.lastByte = int(p[n-1])
				b.lastRuneSize = -1
			}
			return n, b.readErr()
		}
		b.r, b.w = 0, 0
		// read buffer data
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
			panic("non-positive")
		}
		if n == 0 {
			return 0, b.readErr()
		}
		// growth offset
		b.w += n
	}
	// 存在偏移量，copy 缓冲区中这部分偏移量到 p 中
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	// update index
	b.lastByte = int(b.buf[b.r-1])
	b.lastRuneSize = -1
	return n, nil
}

type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return NewWriterSize(w, defaultBufSize)
}

func NewWriterSize(w io.Writer, size int) *Writer {
	b, ok := w.(*Writer)
	if ok && len(b.buf) >= size {
		return b
	}
	if size <= 0 {
		size = defaultBufSize
	}
	return &Writer{
		buf: make([]byte, size),
		wr:  w,
	}
}

func (b *Writer) Buffered() int {
	return b.n
}

// buffer 的可用容量
func (b *Writer) Available() int {
	return len(b.buf) - b.n
}

// Flush kinda complicated so here we just return nil
// you only need to know Flush() will flush data into file
func (b *Writer) Flush() error {
	return nil
}

// 判断buf中可用容量是否可以放下p，如果能放下，直接把p拼接到buf后面，即把内容放到缓冲区
// 如果缓冲区的可用容量不足以放下，且此时缓冲区是空的，直接把p写入文件即可
// 如果缓冲区的可用容量不足以放下，且此时缓冲区有内容，则用p把缓冲区填满，把缓冲区所有内容写入文件，并清空缓冲区
// 判断p的剩余内容大小能否放到缓冲区，如果能放下（此时和步骤1情况一样）则把内容放到缓冲区
// 如果p的剩余内容依旧大于缓冲区，（注意此时缓冲区是空的，情况和步骤2一样）则把p的剩余内容直接写入文件
func (b *Writer) Write(p []byte) (nn int, err error) {
	// 待拷贝数据大于可用缓冲区中的空间且写入过程中没有遇错
	for len(p) > b.Available() && b.err == nil {
		var n int
		// 如果缓存的容量为 0
		if b.Buffered() == 0 {
			// 则将数据写入 p
			n, b.err = b.wr.Write(p)
		} else {
			// 否则 copy n 之后的数据到缓冲区
			n = copy(b.buf[b.n:], p)
			b.n += n
			// 将缓冲区的数据写入文件
			b.Flush()
		}
		// 更新索引
		nn += n
		p = p[n:]
	}
	// 写入过程出错
	if b.err != nil {
		return nn, b.err
	}
	// copy data into buffer
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}

type ReadWriter struct {
	*Reader
	*Writer
}

func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
	return &ReadWriter{r, w}
}
