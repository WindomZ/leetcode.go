package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, twoSum([]int{}, 6), []int(nil))
	assert.Equal(t, twoSum([]int{3}, 6), []int(nil))
	assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	assert.Equal(t, twoSum([]int{3, 2, 4}, 7), []int{0, 2})
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{3, 2, 4}, 6)
	}
}
