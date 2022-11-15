package std

import (
	"path/filepath"
)

type Cmd struct {
	Path string
	Args []string
	Err  error // LookPath error, if any.
}

func Command(name string, arg ...string) *Cmd {
	cmd := &Cmd{
		Path: name,
		Args: append([]string{name}, arg...),
	}
	if filepath.Base(name) == name {
		lp, err := LookPath(name)
		if lp != "" {
			cmd.Path = lp
		}
		if err != nil {
			cmd.Err = err
		}
	}
	return cmd
}

func LookPath(file string) (string, error) {
	return "", nil
}
