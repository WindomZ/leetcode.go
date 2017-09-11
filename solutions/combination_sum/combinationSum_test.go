package combinationsum

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	assert.Equal(t, [][]int{}, combinationSum([]int{}, 0))
	assert.Equal(t, [][]int{{1, 1}}, combinationSum([]int{1}, 2))
	assert.Equal(t, [][]int{{3, 7, 7}}, combinationSum([]int{7, 3}, 17))
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
}

func Benchmark_combinationSum(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			combinationSum([]int{}, 0)
			combinationSum([]int{7, 3}, 17)
			combinationSum([]int{2, 3, 6, 7}, 7)
		}
	})
}
