package twosum

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, []int(nil), twoSum([]int{}, 6))
	assert.Equal(t, []int(nil), twoSum([]int{3}, 6))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 2}, twoSum([]int{3, 2, 4}, 7))
}

func Benchmark_twoSum(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			twoSum([]int{3, 2, 4}, 5)
			twoSum([]int{3, 2, 4}, 6)
			twoSum([]int{3, 2, 4}, 7)
		}
	})
}
