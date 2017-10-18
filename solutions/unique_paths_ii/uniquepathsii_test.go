package uniquepathsii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	assert.Equal(t, 0,
		uniquePathsWithObstacles([][]int{}))
	assert.Equal(t, 2,
		uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}}))
	assert.Equal(t, 3,
		uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}, {0, 0}}))
	assert.Equal(t, 6,
		uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
	assert.Equal(t, 2,
		uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func Benchmark_uniquePathsWithObstacles(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uniquePathsWithObstacles([][]int{})
			uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}})
			uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}, {0, 0}})
			uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})
			uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
		}
	})
}
