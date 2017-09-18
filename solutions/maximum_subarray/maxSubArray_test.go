package maximumsubarray

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	assert.Equal(t, 0, maxSubArray([]int{}))
	assert.Equal(t, 7, maxSubArray([]int{1, 2, 3, -4, 5}))
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 14, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 9, -6}))
	assert.Equal(t, 18, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 9, -6, 10, -12}))
	assert.Equal(t, 10, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, -9, -6, 10, -12}))
}

func Benchmark_maxSubArray(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			maxSubArray([]int{})
			maxSubArray([]int{1, 2, 3, -4, 5})
			maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
			maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 9, -6})
			maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 9, -6, 10, -12})
			maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, -9, -6, 10, -12})
		}
	})
}
