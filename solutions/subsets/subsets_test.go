package subsets

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_subsets(t *testing.T) {
	assert.Equal(t,
		[][]int{}, subsets([]int{}))

	assert.Equal(t,
		[][]int{
			{},
			{1},
			{2},
			{1, 2},
		}, subsets([]int{1, 2}))

	assert.Equal(t,
		[][]int{
			{},
			{2},
			{3},
			{2, 3},
		}, subsets([]int{2, 3}))

	assert.Equal(t,
		[][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}, subsets([]int{1, 2, 3}))

	assert.Equal(t,
		[][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
			{5},
			{1, 5},
			{2, 5},
			{1, 2, 5},
			{3, 5},
			{1, 3, 5},
			{2, 3, 5},
			{1, 2, 3, 5},
		}, subsets([]int{1, 2, 3, 5}))
}

func Benchmark_subsets(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			subsets([]int{})
			subsets([]int{1})
			subsets([]int{2, 3})
			subsets([]int{1, 2, 3})
			subsets([]int{1, 2, 3, 5})
		}
	})
}
