package uniquebinarysearchtrees

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_numTrees(t *testing.T) {
	assert.Equal(t, 1, numTrees(0))
	assert.Equal(t, 1, numTrees(1))
	assert.Equal(t, 2, numTrees(2))
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 14, numTrees(4))
	assert.Equal(t, 42, numTrees(5))
}

func Benchmark_numTrees(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			numTrees(0)
			numTrees(1)
			numTrees(2)
			numTrees(3)
			numTrees(4)
			numTrees(5)
		}
	})
}
