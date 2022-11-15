package std

import (
	"os"
	"syscall"
)

const numSig = 65

func Notify(c chan<- os.Signal, sig ...os.Signal) {
	if c == nil {
		panic("nil channel")
	}

	// update handler

	add := func(n int) {
		if n < 0 {
			return
		}
		// handler check and update
	}
	if len(sig) == 0 {
		for n := 0; n < numSig; n++ {
			add(n)
		}
	} else {
		for _, s := range sig {
			add(signum(s))
		}
	}
}

func signum(sig os.Signal) int {
	switch sig := sig.(type) {
	case syscall.Signal:
		i := int(sig)
		if i < 0 || i >= numSig {
			return -1
		}
		return i
	default:
		return -1
	}
}
