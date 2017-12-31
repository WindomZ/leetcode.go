package removeduplicatesfromsortedarrayii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates([]int{}))
	assert.Equal(t, 1, removeDuplicates([]int{1}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 1}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 1}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 1, 1}))
	assert.Equal(t, 3, removeDuplicates([]int{1, 1, 1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	assert.Equal(t, 6, removeDuplicates([]int{1, 1, 1, 2, 2, 3, 3}))
	assert.Equal(t, 6, removeDuplicates([]int{1, 1, 1, 2, 2, 3, 3, 3}))
}

func Benchmark_removeDuplicates(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			removeDuplicates([]int{})
			removeDuplicates([]int{1, 1, 2})
			removeDuplicates([]int{1, 1, 2, 3, 3})
			removeDuplicates([]int{1, 1, 2, 3, 3, 4, 4, 5})
			removeDuplicates([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7})
			removeDuplicates([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 8, 8})
		}
	})
}
