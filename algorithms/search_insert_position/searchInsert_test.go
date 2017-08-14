package searchinsertposition

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{}, 0))
	assert.Equal(t, 0, searchInsert([]int{4}, 4))
	assert.Equal(t, 1, searchInsert([]int{4, 5}, 5))
	assert.Equal(t, 2, searchInsert([]int{4, 5, 6}, 6))
	assert.Equal(t, 0, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 0))
	assert.Equal(t, 1, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 1))
	assert.Equal(t, 2, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 2))
	assert.Equal(t, 3, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 3))
	assert.Equal(t, 6, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 7))
	assert.Equal(t, 7, searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 8))
}

func Benchmark_searchInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsert([]int{}, 0)
		searchInsert([]int{4}, 4)
		searchInsert([]int{4, 5}, 5)
		searchInsert([]int{4, 5, 6}, 6)
		searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 2)
		searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 3)
		searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 7)
		searchInsert([]int{0, 1, 2, 4, 5, 6, 7}, 8)
	}
}
