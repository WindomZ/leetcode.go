package two_sum

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
	for i := 0; i < b.N; i++ {
		twoSum([]int{3, 2, 4}, 5)
		twoSum([]int{3, 2, 4}, 6)
		twoSum([]int{3, 2, 4}, 7)
	}
}
