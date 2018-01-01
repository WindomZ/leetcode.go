package threesum

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_threeSum(t *testing.T) {
	assert.Equal(t, [][]int{}, threeSum([]int{}))
	assert.Equal(t, [][]int{}, threeSum([]int{1, 2}))
	assert.Equal(t, [][]int{{-4, 1, 3}}, threeSum([]int{1, 3, -4}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0, 0}))
	assert.Equal(t,
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4}),
	)
	assert.Equal(t,
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4, -1, 1}),
	)
}

func Benchmark_threeSum(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			threeSum([]int{})
			threeSum([]int{1, 3, -4})
			threeSum([]int{0, 0, 0, 0})
			threeSum([]int{-1, 0, 1, 2, -1, -4})
		}
	})
}
