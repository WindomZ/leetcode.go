package searchinrotatedsortedarray

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_search(t *testing.T) {
	assert.Equal(t, -1, search([]int{}, 0))
	assert.Equal(t, 0, search([]int{4}, 4))
	assert.Equal(t, 1, search([]int{4, 5}, 5))
	assert.Equal(t, 2, search([]int{4, 5, 6}, 6))
	assert.Equal(t, 0, search([]int{4, 5, 6, 7, 0, 1, 2}, 4))
	assert.Equal(t, 1, search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, 6, search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}

func Benchmark_search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search([]int{}, 0)
		search([]int{4}, 4)
		search([]int{4, 5}, 5)
		search([]int{4, 5, 6}, 6)
		search([]int{4, 5, 6, 7, 0, 1, 2}, 4)
		search([]int{4, 5, 6, 7, 0, 1, 2}, 5)
		search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
		search([]int{4, 5, 6, 7, 0, 1, 2}, 2)
	}
}
