package permutationsii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_permuteUnique(t *testing.T) {
	assert.Equal(t, [][]int{}, permuteUnique([]int{}))
	assert.Equal(t, [][]int{
		{1, 1, 1},
	}, permuteUnique([]int{1, 1, 1}))
	assert.Equal(t, [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
	}, permuteUnique([]int{1, 2, 3}))
	assert.Equal(t, [][]int{
		{1, 1, 2, 2},
		{1, 2, 1, 2},
		{1, 2, 2, 1},
		{2, 1, 1, 2},
		{2, 1, 2, 1},
		{2, 2, 1, 1},
	}, permuteUnique([]int{2, 2, 1, 1}))
	assert.Equal(t, [][]int{
		{1, 1, 2, 3},
		{1, 1, 3, 2},
		{1, 2, 1, 3},
		{1, 2, 3, 1},
		{1, 3, 2, 1},
		{1, 3, 1, 2},
		{2, 1, 1, 3},
		{2, 1, 3, 1},
		{2, 3, 1, 1},
		{3, 1, 2, 1},
		{3, 1, 1, 2},
		{3, 2, 1, 1},
	}, permuteUnique([]int{1, 1, 2, 3}))

	//permuteUnique([]int{-1, 2, 0, -1, 1, 0, 1})
}

func Benchmark_permuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permuteUnique([]int{})
		permuteUnique([]int{1, 1, 2})
		permuteUnique([]int{1, 2, 3})
	}
}
