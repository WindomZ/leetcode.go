package minimumpathsum

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_minPathSum(t *testing.T) {
	assert.Empty(t, minPathSum([][]int{{0}}))
	assert.Equal(t, 4, minPathSum([][]int{
		{1, 2},
		{2, 1},
	}))
	assert.Equal(t, 4, minPathSum([][]int{
		{1, 2},
		{3, 1},
	}))
	assert.Equal(t, 7, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	assert.Equal(t, 10, minPathSum([][]int{
		{1, 3, 1, 2},
		{1, 5, 2, 2},
		{4, 2, 1, 2},
	}))
	assert.Equal(t, 11, minPathSum([][]int{
		{1, 3, 1, 2},
		{1, 5, 2, 2},
		{4, 2, 1, 2},
		{1, 0, 4, 1},
	}))
}

func Benchmark_permute(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			minPathSum([][]int{{0}})
			minPathSum([][]int{
				{1, 2},
				{2, 1},
			})
			minPathSum([][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			})
			minPathSum([][]int{
				{1, 3, 1, 2},
				{1, 5, 2, 2},
				{4, 2, 1, 2},
			})
			minPathSum([][]int{
				{1, 3, 1, 2},
				{1, 5, 2, 2},
				{4, 2, 1, 2},
				{1, 0, 4, 1},
			})
		}
	})
}
