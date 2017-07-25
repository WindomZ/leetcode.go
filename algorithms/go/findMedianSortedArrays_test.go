package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{1}, []int{}))
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{}, []int{1}))
	assert.Equal(t, float64(1.5), findMedianSortedArrays([]int{1, 2}, []int{}))
	assert.Equal(t, float64(1.5), findMedianSortedArrays([]int{}, []int{1, 2}))

	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1}, []int{2, 3}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{2}, []int{1, 3}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{3}, []int{1, 2}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 2}, []int{3}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{2, 3}, []int{1}))

	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1}, []int{2, 3, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{2}, []int{1, 3, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{3}, []int{1, 2, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{4}, []int{1, 2, 3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 4}, []int{2, 3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{2, 3}, []int{1, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{2, 4}, []int{1, 3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{3, 4}, []int{1, 2}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2, 3}, []int{4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2, 4}, []int{3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 3, 4}, []int{2}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{2, 3, 4}, []int{1}))

	assert.Equal(t, float64(3), findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5}))
	assert.Equal(t, float64(3), findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5}))
	assert.Equal(t, float64(3), findMedianSortedArrays([]int{2, 3, 4, 5}, []int{1}))

	defer func() {
		assert.NotEmpty(t, recover())
	}()
	findMedianSortedArrays([]int{}, []int{})
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	}
}
