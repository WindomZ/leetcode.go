package sqrtx

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(2))
	assert.Equal(t, 1, mySqrt(3))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 3, mySqrt(9))
	assert.Equal(t, 3, mySqrt(15))
	assert.Equal(t, 4, mySqrt(16))
	assert.Equal(t, 4, mySqrt(24))
	assert.Equal(t, 5, mySqrt(25))
}

func Benchmark_mySqrt(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mySqrt(0)
			mySqrt(1)
			mySqrt(3)
			mySqrt(8)
			mySqrt(9)
			mySqrt(15)
			mySqrt(16)
			mySqrt(25)
		}
	})
}
