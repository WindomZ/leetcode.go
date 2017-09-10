package rotateimage

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_rotate(t *testing.T) {
	var matrix [][]int

	matrix = [][]int{}
	rotate(matrix)
	assert.Equal(t, [][]int{}, matrix)

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	assert.Equal(t, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	assert.Equal(t, [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}, matrix)
}

func Benchmark_rotate(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rotate([][]int{})
			rotate([][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			})
			rotate([][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			})
		}
	})
}
