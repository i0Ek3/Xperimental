package std

import (
	"io"
	"io/fs"
	"syscall"
)

// fd
type File struct{}

var ErrInvalid = fs.ErrInvalid

const (
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
)

func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}

type FileMode = fs.FileMode

func OpenFile(name string, flag int, perm FileMode) (f *File, err error) { return }

func (f *File) Read(b []byte) (n int, err error) {
	if err := f.checkValid("read"); err != nil {
		return 0, err
	}
	n, e := f.read(b)
	return n, f.wrapErr("read", e)
}

func (f *File) read(b []byte) (n int, err error) { return }

func (f *File) checkValid(op string) error {
	if f == nil {
		return ErrInvalid
	}
	return nil
}

func (f *File) wrapErr(op string, err error) error {
	if err == nil || err == io.EOF {
		return err
	}
	return nil
}
