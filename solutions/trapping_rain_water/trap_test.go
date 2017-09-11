package trappingrainwater

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_trap(t *testing.T) {
	assert.Equal(t, 0, trap([]int{}))
	assert.Equal(t, 3, trap([]int{4, 1, 0, 2, 1, 0}))
	assert.Equal(t, 8, trap([]int{4, 4, 2, 2, 1, 3, 5}))
	assert.Equal(t, 8, trap([]int{4, 4, 2, 2, 1, 3, 5, 5, 5}))
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 16, trap([]int{2, 1, 6, 6, 1, 1, 1, 3, 5, 1, 2, 1}))
}

func Benchmark_trap(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			trap([]int{})
			trap([]int{4, 1, 0, 2, 1, 0})
			trap([]int{4, 4, 2, 2, 1, 3, 5})
			trap([]int{4, 4, 2, 2, 1, 3, 5, 5, 5})
			trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
			trap([]int{2, 1, 6, 6, 1, 1, 1, 3, 5, 1, 2, 1})
		}
	})
}
