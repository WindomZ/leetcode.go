package powxn

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_myPow(t *testing.T) {
	assert.Equal(t, float64(1), myPow(0, 0))
	assert.Equal(t, float64(1), myPow(1, -1))
	assert.Equal(t, float64(1), myPow(-1, math.MinInt32))
	assert.Equal(t, float64(1), myPow(1, math.MinInt32))
	assert.Equal(t, float64(0), myPow(-2, math.MinInt32))
	assert.Equal(t, float64(0), myPow(2, math.MinInt32))
	assert.Equal(t, float64(1), myPow(1, 2))
	assert.Equal(t, float64(1), myPow(1, 3))
	assert.Equal(t, float64(1), myPow(1, 4))
	assert.Equal(t, float64(0.25), myPow(2, -2))
	assert.Equal(t, float64(4), myPow(2, 2))
	assert.Equal(t, float64(0.125), myPow(2, -3))
	assert.Equal(t, float64(8), myPow(2, 3))
}

func Benchmark_myPow(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			myPow(0, 0)
			myPow(1, -1)
			myPow(1, math.MinInt32)
			myPow(1, 2)
			myPow(1, 3)
			myPow(1, 4)
			myPow(2, -2)
			myPow(2, 2)
			myPow(2, -3)
			myPow(2, 3)
			myPow(10, -10)
			myPow(10, 10)
		}
	})
}
