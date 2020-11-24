package msort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int) bool
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
	}
}

type IntArray [10]int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) bool {
	p[i], p[j] = p[j], p[i]
	return true
}
