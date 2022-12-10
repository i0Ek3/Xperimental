package lock_chan

import "testing"

func BenchmarkChanMutex(b *testing.B) {
	c := NewChanMutex()
	for i := 0; i < b.N; i++ {
		c.Add(i)
	}
}

func BenchmarkCommonMutex(b *testing.B) {
	c := NewCommonMutex()
	for i := 0; i < b.N; i++ {
		c.Add(i)
	}
}
