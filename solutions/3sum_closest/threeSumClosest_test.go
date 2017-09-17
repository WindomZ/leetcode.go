package threesumclosest

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	assert.Equal(t, 0, threeSumClosest([]int{}, 0))
	assert.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 5))
	assert.Equal(t, -4, threeSumClosest([]int{-1, 2, 1, -4, 5, 8}, -5))
	assert.Equal(t, 10, threeSumClosest([]int{-1, 2, 1, -4, 5, -8, 9}, 11))
	assert.Equal(t, 28, threeSumClosest([]int{-1, 2, 1, -4, 2, -8, 9, 10, 9}, 25))
	assert.Equal(t, -19, threeSumClosest([]int{-1, 2, 1, -4, 5, -8, -7, 9, 10, 9}, -20))
	assert.Equal(t, 0, threeSumClosest([]int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}, 0))
}

func Benchmark_threeSumClosest(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			threeSumClosest([]int{}, 0)
			threeSumClosest([]int{0, 0, 0}, 1)
			threeSumClosest([]int{-1, 2, 1, -4}, 5)
			threeSumClosest([]int{-1, 2, 1, -4, 5, 8}, -5)
			threeSumClosest([]int{-1, 2, 1, -4, 5, -8, 9}, 11)
			threeSumClosest([]int{-1, 2, 1, -4, 2, -8, 9, 10, 9}, 25)
			threeSumClosest([]int{-1, 2, 1, -4, 5, -8, -7, 9, 10, 9}, -20)
			threeSumClosest([]int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}, 0)
		}
	})
}
