package subsetsii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	assert.Equal(t, [][]int{}, subsetsWithDup([]int{}))
	assert.Equal(t, [][]int{
		{},
		{1},
	}, subsetsWithDup([]int{1}))
	assert.Equal(t, [][]int{
		{},
		{1},
		{2},
		{1, 2},
	}, subsetsWithDup([]int{1, 2}))
	assert.Equal(t, [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{2, 2},
		{1, 2, 2},
	}, subsetsWithDup([]int{1, 2, 2}))
	assert.Equal(t, [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{2, 2},
		{1, 2, 2},
		{2, 2, 2},
		{1, 2, 2, 2},
	}, subsetsWithDup([]int{1, 2, 2, 2}))
	assert.Equal(t, [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{2, 2},
		{1, 2, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
		{2, 2, 3},
		{1, 2, 2, 3},
	}, subsetsWithDup([]int{1, 3, 2, 2}))
}

func Benchmark_subsetsWithDup(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			subsetsWithDup([]int{})
			subsetsWithDup([]int{1})
			subsetsWithDup([]int{1, 2})
			subsetsWithDup([]int{1, 2, 2})
			subsetsWithDup([]int{1, 2, 2, 3})
		}
	})
}
