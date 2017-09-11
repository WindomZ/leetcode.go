package combinationsumii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
	assert.Equal(t, [][]int{}, combinationSum2([]int{}, 0))
	assert.Equal(t, [][]int{{2}}, combinationSum2([]int{1, 2}, 2))
	assert.Equal(t, [][]int{{7}}, combinationSum2([]int{2, 3, 6, 7}, 7))
	assert.Equal(t, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func Benchmark_combinationSum2(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			combinationSum2([]int{}, 0)
			combinationSum2([]int{2, 3, 6, 7}, 7)
			combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
		}
	})
}
