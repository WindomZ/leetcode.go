package removeelement

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_removeElement(t *testing.T) {
	assert.Equal(t, 0, removeElement([]int{}, 0))
	assert.Equal(t, 1, removeElement([]int{1, 1, 2}, 1))
	assert.Equal(t, 3, removeElement([]int{1, 1, 2, 3, 3}, 3))
	assert.Equal(t, 6, removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5}, 4))
	assert.Equal(t, 9, removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7}, 6))
	assert.Equal(t, 11, removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 8, 8}, 8))
}

func Benchmark_removeElement(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			removeElement([]int{}, 0)
			removeElement([]int{1, 1, 2}, 1)
			removeElement([]int{1, 1, 2, 3, 3}, 3)
			removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5}, 4)
			removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7}, 6)
			removeElement([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 8, 8}, 8)
		}
	})
}
