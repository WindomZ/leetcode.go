package permutations

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_permute(t *testing.T) {
	assert.Equal(t, [][]int{}, permute([]int{}))
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
	}, permute([]int{1, 2, 3}))
	assert.Equal(t, [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 3, 2},
		{1, 4, 2, 3},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 3, 1},
		{2, 4, 1, 3},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 2, 3, 1},
		{4, 2, 1, 3},
		{4, 3, 2, 1},
		{4, 3, 1, 2},
		{4, 1, 3, 2},
		{4, 1, 2, 3},
	}, permute([]int{1, 2, 3, 4}))
}

func Benchmark_permute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permute([]int{})
		permute([]int{1, 2, 3})
		permute([]int{1, 2, 3, 4})
	}
}
