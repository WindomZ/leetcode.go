package firstmissingpositive

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_firstMissingPositive(t *testing.T) {
	assert.Equal(t, 1, firstMissingPositive([]int{}))
	assert.Equal(t, 3, firstMissingPositive([]int{1, 2, 0}))
	assert.Equal(t, 2, firstMissingPositive([]int{3, 4, -1, 1}))
	assert.Equal(t, 6, firstMissingPositive([]int{3, 4, -1, 1, 2, 5}))
	assert.Equal(t, 6, firstMissingPositive([]int{3, 4, -1, 1, 2, 5, 8, 9}))
	assert.Equal(t, 7, firstMissingPositive([]int{6, 4, -1, -2, 2, 5, 3, 1}))
}

func Benchmark_firstMissingPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstMissingPositive([]int{})
		firstMissingPositive([]int{1, 2, 0})
		firstMissingPositive([]int{3, 4, -1, 1})
		firstMissingPositive([]int{3, 4, -1, 1, 2, 5})
		firstMissingPositive([]int{3, 4, -1, 1, 2, 5, 8, 9})
		firstMissingPositive([]int{6, 4, -1, -2, 2, 5, 3, 1})
	}
}
