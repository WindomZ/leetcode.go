package searchforarange

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_searchRange(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
	assert.Equal(t, []int{0, 0}, searchRange([]int{4}, 4))
	assert.Equal(t, []int{1, 1}, searchRange([]int{4, 5}, 5))
	assert.Equal(t, []int{2, 2}, searchRange([]int{4, 5, 6}, 6))
	assert.Equal(t, []int{2, 3}, searchRange([]int{4, 5, 6, 6, 7}, 6))
	assert.Equal(t, []int{4, 5}, searchRange([]int{4, 5, 6, 6, 7, 7, 8}, 7))
	assert.Equal(t, []int{4, 6}, searchRange([]int{4, 5, 6, 6, 7, 7, 7, 8}, 7))
	assert.Equal(t, []int{7, 8}, searchRange([]int{4, 5, 6, 6, 7, 7, 7, 8, 8}, 8))
}

func Benchmark_searchRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRange([]int{}, 0)
		searchRange([]int{4}, 4)
		searchRange([]int{4, 5}, 5)
		searchRange([]int{4, 5, 6}, 6)
		searchRange([]int{4, 5, 6, 6, 7}, 6)
		searchRange([]int{4, 5, 6, 6, 7, 7, 8}, 7)
		searchRange([]int{4, 5, 6, 6, 7, 7, 7, 8}, 7)
		searchRange([]int{4, 5, 6, 6, 7, 7, 7, 8, 8}, 8)
	}
}
