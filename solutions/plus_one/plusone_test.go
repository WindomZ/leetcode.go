package plusone

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_plusOne(t *testing.T) {
	assert.Equal(t, []int{1}, plusOne([]int{}))
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t, []int{1, 1}, plusOne([]int{1, 0}))
	assert.Equal(t, []int{1, 6}, plusOne([]int{1, 5}))
	assert.Equal(t, []int{1, 0, 0}, plusOne([]int{9, 9}))
}

func Benchmark_plusOne(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			plusOne([]int{0})
			plusOne([]int{9})
			plusOne([]int{1, 9})
			plusOne([]int{9, 9})
			plusOne([]int{1, 9, 9})
			plusOne([]int{9, 9, 9})
		}
	})
}
