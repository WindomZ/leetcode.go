package sortcolors

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_sortColors(t *testing.T) {
	nums := []int{0}
	sortColors(nums)
	assert.Equal(t, []int{0}, nums)

	nums = []int{0, 2, 1}
	sortColors(nums)
	assert.Equal(t, []int{0, 1, 2}, nums)

	nums = []int{0, 2, 1, 2}
	sortColors(nums)
	assert.Equal(t, []int{0, 1, 2, 2}, nums)

	nums = []int{0, 2, 1, 2, 0, 2, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 0, 0, 1, 2, 2, 2}, nums)

	nums = []int{0, 2, 1, 2, 0, 2, 0, 1, 2, 1, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2}, nums)

	nums = []int{0, 2, 1, 2, 0, 2, 0, 1, 2, 1, 0, 1, 2, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2}, nums)
}

func Benchmark_sortColors(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sortColors([]int{0})
			sortColors([]int{0, 2, 1})
			sortColors([]int{0, 2, 1, 2})
			sortColors([]int{0, 2, 1, 2, 0, 2, 0})
			sortColors([]int{0, 2, 1, 2, 0, 2, 0, 1, 2, 1, 0})
			sortColors([]int{0, 2, 1, 2, 0, 2, 0, 1, 2, 1, 0, 1, 2, 0})
		}
	})
}
