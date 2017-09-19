package jumpgame

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_canJump(t *testing.T) {
	assert.Equal(t, true, canJump([]int{}))
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))
	assert.Equal(t, true, canJump([]int{3, 2, 1, 2, 4, 5}))
	assert.Equal(t, true, canJump([]int{1, 2, 1, 2, 1, 1, 2, 1}))
	assert.Equal(t, true, canJump([]int{4, 2, 1, 2, 1, 5, 2, 1, 2}))
}

func Benchmark_canJump(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			canJump([]int{})
			canJump([]int{2, 3, 1, 1, 4})
			canJump([]int{3, 2, 1, 0, 4})
			canJump([]int{3, 2, 1, 2, 4, 5})
			canJump([]int{1, 2, 1, 2, 1, 1, 2, 1})
			canJump([]int{4, 2, 1, 2, 1, 5, 2, 1, 2})
		}
	})
}
