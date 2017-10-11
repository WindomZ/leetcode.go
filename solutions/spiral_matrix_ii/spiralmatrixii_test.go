package spiralmatrixii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_generateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{},
		generateMatrix(0))
	assert.Equal(t, [][]int{{1}},
		generateMatrix(1))
	assert.Equal(t, [][]int{{1, 2}, {4, 3}},
		generateMatrix(2))
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		generateMatrix(3))
	assert.Equal(t, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		generateMatrix(4))
}

func Benchmark_generateMatrix(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generateMatrix(1)
			generateMatrix(2)
			generateMatrix(3)
		}
	})
}
