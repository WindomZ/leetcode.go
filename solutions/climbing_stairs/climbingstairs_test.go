package climbingstairs

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, 1, climbStairs(1))
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 5, climbStairs(4))
	assert.Equal(t, 8, climbStairs(5))
}

func Benchmark_climbStairs(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			climbStairs(1)
			climbStairs(2)
			climbStairs(3)
			climbStairs(4)
			climbStairs(5)
			climbStairs(6)
			climbStairs(7)
			climbStairs(8)
			climbStairs(9)
		}
	})
}
