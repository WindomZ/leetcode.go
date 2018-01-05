package mergesortedarray

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_merge(t *testing.T) {
	num := []int{1, 0}
	merge(num, 1, []int{2}, 1)
	assert.Equal(t, []int{1, 2}, num)

	num = []int{1, 2, 0, 0}
	merge(num, 2, []int{2, 3}, 2)
	assert.Equal(t, []int{1, 2, 2, 3}, num)

	num = []int{1, 2, 4, 0, 0}
	merge(num, 3, []int{2, 3}, 2)
	assert.Equal(t, []int{1, 2, 2, 3, 4}, num)

	num = []int{1, 2, 4, 0}
	merge(num, 2, []int{2, 3}, 2)
	assert.Equal(t, []int{1, 2, 2, 3}, num)

	num = []int{3, 4, 5, 0, 0, 0}
	merge(num, 3, []int{1, 2, 3}, 3)
	assert.Equal(t, []int{1, 2, 3, 3, 4, 5}, num)

	num = []int{4, 5, 6, 0, 0, 0}
	merge(num, 3, []int{1, 2, 3}, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, num)
}

func Benchmark_merge(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			merge([]int{1, 0}, 1, []int{2}, 1)
			merge([]int{1, 2, 0, 0}, 2, []int{2, 3}, 2)
			merge([]int{1, 2, 4, 0, 0}, 3, []int{2, 3}, 2)
			merge([]int{1, 2, 4, 0}, 2, []int{2, 3}, 2)
			merge([]int{3, 4, 5, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
			merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
		}
	})
}
