package spiralmatrix

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_spiralOrder(t *testing.T) {
	assert.Equal(t, []int{},
		spiralOrder([][]int{}))
	assert.Equal(t, []int{1, 2, 5, 8, 7, 4},
		spiralOrder([][]int{
			{1, 2},
			{4, 5},
			{7, 8},
		}))
	assert.Equal(t, []int{1, 2, 3, 6, 5, 4},
		spiralOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
		}))
	assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		spiralOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}))
	assert.Equal(t, []int{1, 2, 3, 2, 1, 4, 5, 6, 7, 6, 5, 4, 5, 6, 5},
		spiralOrder([][]int{
			{1, 2, 3, 2, 1},
			{4, 5, 6, 5, 4},
			{5, 6, 7, 6, 5},
		}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 3, 4, 5},
		spiralOrder([][]int{
			{1, 2, 3},
			{2, 3, 4},
			{3, 4, 5},
			{4, 5, 6},
			{5, 6, 7},
		}))
	assert.Equal(t, []int{1, 2, 3, 2, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 3, 4, 3, 4, 5, 6, 5, 4, 5},
		spiralOrder([][]int{
			{1, 2, 3, 2, 1},
			{2, 3, 4, 3, 2},
			{3, 4, 5, 4, 3},
			{4, 5, 6, 5, 4},
			{5, 6, 7, 6, 5},
		}))
}

func Benchmark_spiralOrder(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spiralOrder([][]int{})
			spiralOrder([][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			})
			spiralOrder([][]int{
				{1, 2, 3, 2, 1},
				{2, 3, 4, 3, 2},
				{3, 4, 5, 4, 3},
				{4, 5, 6, 5, 4},
				{5, 6, 7, 6, 5},
			})
		}
	})
}
