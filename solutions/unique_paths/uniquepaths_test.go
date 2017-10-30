package uniquepaths

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_uniquePaths(t *testing.T) {
	assert.Equal(t, 1, uniquePaths(1, 1))
	assert.Equal(t, 1, uniquePaths(1, 2))
	assert.Equal(t, 2, uniquePaths(2, 2))
	assert.Equal(t, 3, uniquePaths(2, 3))
	assert.Equal(t, 4, uniquePaths(2, 4))
	assert.Equal(t, 5, uniquePaths(2, 5))
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 6, uniquePaths(3, 3))
	assert.Equal(t, 10, uniquePaths(3, 4))
	assert.Equal(t, 15, uniquePaths(3, 5))
	assert.Equal(t, 924, uniquePaths(7, 7))
	assert.Equal(t, 48620, uniquePaths(10, 10))
}

func Benchmark_uniquePaths(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uniquePaths(1, 1)
			uniquePaths(1, 2)
			uniquePaths(2, 2)
			uniquePaths(2, 3)
			uniquePaths(2, 4)
			uniquePaths(2, 5)
			uniquePaths(3, 2)
			uniquePaths(3, 3)
			uniquePaths(3, 4)
			uniquePaths(3, 5)
		}
	})
}
