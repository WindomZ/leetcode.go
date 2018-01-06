package graycode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_grayCode(t *testing.T) {
	assert.Equal(t, []int{0}, grayCode(0))
	assert.Equal(t, []int{0, 1}, grayCode(1))
	assert.Equal(t, []int{0, 1, 3, 2}, grayCode(2))
	assert.Equal(t, []int{0, 1, 3, 2, 6, 7, 5, 4}, grayCode(3))
	assert.Equal(t, []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}, grayCode(4))
}

func Benchmark_grayCode(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grayCode(0)
			grayCode(1)
			grayCode(2)
			grayCode(3)
			grayCode(4)
			grayCode(5)
		}
	})
}
