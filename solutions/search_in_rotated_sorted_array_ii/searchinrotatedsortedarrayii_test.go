package searchinrotatedsortedarrayii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_search(t *testing.T) {
	assert.Equal(t, false, search([]int{}, 0))
	assert.Equal(t, false, search([]int{1}, 0))
	assert.Equal(t, true, search([]int{1}, 1))
	assert.Equal(t, true, search([]int{0, 1}, 1))
	assert.Equal(t, true, search([]int{0, 1, 2}, 1))
	assert.Equal(t, true, search([]int{0, 1, 2, 3}, 3))
	assert.Equal(t, true, search([]int{0, 1, 2, 2, 3}, 2))
	assert.Equal(t, true, search([]int{0, 1, 2, 2, 3, 3}, 2))
	assert.Equal(t, true, search([]int{0, 1, 2, 2, 3, 3, 3}, 2))
	assert.Equal(t, true, search([]int{0, 1, 2, 2, 3, 3, 3, 4}, 3))
	assert.Equal(t, true, search([]int{0, 1, 2, 2, 3, 3, 3, 3, 3}, 0))
	assert.Equal(t, true, search([]int{3, 3, 4, 4, 0, 1, 1, 2, 2}, 3))
	assert.Equal(t, true, search([]int{1, 1, 2, 2, 3, 3, 4, 4, 0, 0}, 0))
}

func Benchmark_search(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			search([]int{}, 0)
			search([]int{1}, 0)
			search([]int{1}, 1)
			search([]int{0, 1}, 1)
			search([]int{0, 1, 2}, 1)
			search([]int{0, 1, 2, 3}, 3)
			search([]int{0, 1, 2, 2, 3}, 2)
			search([]int{0, 1, 2, 2, 3, 3}, 2)
			search([]int{0, 1, 2, 2, 3, 3, 3}, 2)
			search([]int{0, 1, 2, 2, 3, 3, 3, 4}, 3)
			search([]int{3, 3, 4, 4, 0, 1, 1, 2, 2}, 3)
			search([]int{1, 1, 2, 2, 3, 3, 4, 4, 0, 0}, 0)
		}
	})
}
