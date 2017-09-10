package foursum

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_fourSum(t *testing.T) {
	assert.Equal(t, [][]int{}, fourSum([]int{}, 0))
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	assert.Equal(t, [][]int{
		{-3, -2, 2, 3},
		{-3, -1, 1, 3},
		{-3, 0, 0, 3},
		{-3, 0, 1, 2},
		{-2, -1, 0, 3},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2, -3, 3}, 0))
}

func Benchmark_fourSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum([]int{}, 0)
		fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
		fourSum([]int{1, 0, -1, 0, -2, 2, -3, 3}, 0)
	}
}
