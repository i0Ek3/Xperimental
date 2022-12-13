package lock_chan

import "testing"

func TestChanMutex(t *testing.T) {
	c := NewChanMutex()
	got := c.Add()
	if got != 4950 {
		t.Errorf("wrong result, expect %v", got)
	}
}

func BenchmarkChanMutex(b *testing.B) {
	c := NewChanMutex()
	for i := 0; i < b.N; i++ {
		_ = c.Add()
	}
}
func TestCommonMutex(t *testing.T) {
	c := NewCommonMutex()
	got := c.Add()
	if got != 4950 {
		t.Errorf("wrong result, expect %v", got)
	}
}

func BenchmarkCommonMutex(b *testing.B) {
	c := NewCommonMutex()
	for i := 0; i < b.N; i++ {
		_ = c.Add()
	}
}
