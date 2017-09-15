package nextpermutation

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	var r []int

	r = []int{}
	nextPermutation(r)
	assert.Equal(t, []int{}, r)

	r = []int{1, 2, 3}
	nextPermutation(r)
	assert.Equal(t, []int{1, 3, 2}, r)

	r = []int{3, 2, 1}
	nextPermutation(r)
	assert.Equal(t, []int{1, 2, 3}, r)

	r = []int{1, 1, 5}
	nextPermutation(r)
	assert.Equal(t, []int{1, 5, 1}, r)

	r = []int{1, 2, 5, 1}
	nextPermutation(r)
	assert.Equal(t, []int{1, 5, 1, 2}, r)
}

func Benchmark_nextPermutation(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nextPermutation([]int{})
			nextPermutation([]int{1, 2, 3})
			nextPermutation([]int{3, 2, 1})
			nextPermutation([]int{1, 1, 5})
		}
	})
}
