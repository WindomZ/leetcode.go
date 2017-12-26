package searcha2dmatrix

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	assert.Equal(t, false, searchMatrix([][]int{}, 0))
	assert.Equal(t, false, searchMatrix([][]int{{}}, 1))
	assert.Equal(t, true, searchMatrix([][]int{
		{0},
	}, 0))
	assert.Equal(t, false, searchMatrix([][]int{
		{0},
	}, 1))
	assert.Equal(t, true, searchMatrix([][]int{
		{0, 1},
		{2, 3},
	}, 2))
	assert.Equal(t, false, searchMatrix([][]int{
		{0, 1},
		{2, 3},
	}, 4))
	assert.Equal(t, true, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 50))
	assert.Equal(t, false, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 12))
}

func Benchmark_searchMatrix(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			searchMatrix([][]int{
				{0},
			}, 0)
			searchMatrix([][]int{
				{0},
			}, 1)
			searchMatrix([][]int{
				{0, 1},
				{2, 3},
			}, 2)
			searchMatrix([][]int{
				{0, 1},
				{2, 3},
			}, 4)
			searchMatrix([][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 50)
			searchMatrix([][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 12)
		}
	})
}
