package combinations

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_combine(t *testing.T) {
	assert.Equal(t, [][]int{}, combine(1, 0))
	assert.Equal(t, [][]int{
		{1},
	}, combine(1, 1))
	assert.Equal(t, [][]int{
		{1, 2},
	}, combine(2, 2))
	assert.Equal(t, [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}, combine(3, 2))
	assert.Equal(t, [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}, combine(4, 2))
}

func Benchmark_combine(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			combine(1, 0)
			combine(1, 1)
			combine(2, 2)
			combine(3, 2)
			combine(4, 2)
			combine(4, 3)
		}
	})
}
