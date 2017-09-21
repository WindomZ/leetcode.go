package mergeintervals

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_merge(t *testing.T) {
	assert.Equal(t, []Interval{}, merge([]Interval{}))
	assert.Equal(t,
		[]Interval{
			{1, 4},
		},
		merge([]Interval{
			{1, 4},
			{2, 3},
		}))
	assert.Equal(t,
		[]Interval{
			{1, 6},
			{8, 10},
			{15, 18},
		},
		merge([]Interval{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}))
	assert.Equal(t,
		[]Interval{
			{1, 6},
			{8, 10},
			{15, 18},
		},
		merge([]Interval{
			{2, 6},
			{15, 18},
			{8, 10},
			{1, 3},
		}))
	assert.Equal(t,
		[]Interval{
			{1, 3},
			{4, 7},
		},
		merge([]Interval{
			{2, 3},
			{2, 2},
			{3, 3},
			{1, 3},
			{5, 7},
			{2, 2},
			{4, 6},
		}))
}

func Benchmark_merge(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			merge([]Interval{})
			merge([]Interval{
				{1, 4},
				{2, 3},
			})
			merge([]Interval{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			})
			merge([]Interval{
				{2, 6},
				{15, 18},
				{8, 10},
				{1, 3},
			})
			merge([]Interval{
				{2, 3},
				{2, 2},
				{3, 3},
				{1, 3},
				{5, 7},
				{2, 2},
				{4, 6},
			})
		}
	})
}
