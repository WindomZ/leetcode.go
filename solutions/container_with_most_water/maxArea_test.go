package containerwithmostwater

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_maxArea(t *testing.T) {
	assert.Equal(t, 0, maxArea([]int{}))
	assert.Equal(t, 0, maxArea([]int{1}))
	assert.Equal(t, 1, maxArea([]int{1, 2}))
	assert.Equal(t, 2, maxArea([]int{1, 2, 3}))
	assert.Equal(t, 6, maxArea([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 6, maxArea([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 12, maxArea([]int{5, 1, 3, 4, 2}))
	assert.Equal(t, 16, maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 30, maxArea([]int{8, 1, 7, 2, 6, 3, 5, 4}))
}

func Benchmark_maxArea(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			maxArea([]int{})
			maxArea([]int{1, 2, 3})
			maxArea([]int{5, 4, 3, 2, 1})
			maxArea([]int{1, 2, 3, 4, 5, 6, 7})
			maxArea([]int{8, 1, 7, 2, 6, 3, 5, 4})
		}
	})
}
