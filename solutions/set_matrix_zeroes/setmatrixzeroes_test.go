package setmatrixzeroes

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{0}}
	setZeroes(matrix)
	assert.Equal(t, [][]int{{0}}, matrix)

	matrix = [][]int{
		{0, 1},
		{1, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{0, 0},
		{0, 1},
	}, matrix)

	matrix = [][]int{
		{1, 2},
		{2, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{1, 2},
		{2, 1},
	}, matrix)

	matrix = [][]int{
		{1, 2, 0},
		{2, 1, 2},
		{1, 2, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{0, 0, 0},
		{2, 1, 0},
		{1, 2, 0},
	}, matrix)

	matrix = [][]int{
		{1, 2, 1},
		{2, 0, 2},
		{1, 2, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}, matrix)
}

func Benchmark_setZeroes(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			setZeroes([][]int{{0}})
			setZeroes([][]int{
				{0, 1},
				{1, 1},
			})
			setZeroes([][]int{
				{1, 2},
				{2, 1},
			})
			setZeroes([][]int{
				{1, 2, 0},
				{2, 1, 2},
				{1, 2, 1},
			})
			setZeroes([][]int{
				{1, 2, 1},
				{2, 0, 2},
				{1, 2, 1},
			})
			setZeroes([][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 1},
				{1, 2, 2, 1},
				{1, 1, 1, 1},
			})
		}
	})
}
