package jumpgameii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_jump(t *testing.T) {
	assert.Equal(t, 0, jump([]int{}))
	assert.Equal(t, 0, jump([]int{1}))
	assert.Equal(t, 0, jump([]int{1, 1, 0, 0}))
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 3, jump([]int{1, 2, 4, 3, 6}))
	assert.Equal(t, 1, jump([]int{4, 2, 1, 3, 6}))
}

func Benchmark_jump(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jump([]int{})
			jump([]int{1})
			jump([]int{1, 1, 0, 0})
			jump([]int{2, 3, 1, 1, 4})
			jump([]int{1, 2, 4, 3, 6})
			jump([]int{4, 2, 1, 3, 6})
		}
	})
}
