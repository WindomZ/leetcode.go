package permutationsequence

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	assert.Equal(t, "", getPermutation(0, 0))
	assert.Equal(t, "1", getPermutation(1, 1))
	assert.Equal(t, "123", getPermutation(3, 1))
	assert.Equal(t, "132", getPermutation(3, 2))
	assert.Equal(t, "213", getPermutation(3, 3))
	assert.Equal(t, "231", getPermutation(3, 4))
	assert.Equal(t, "312", getPermutation(3, 5))
	assert.Equal(t, "321", getPermutation(3, 6))
}

func Benchmark_getPermutation(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getPermutation(0, 0)
			getPermutation(1, 1)
			getPermutation(3, 1)
			getPermutation(3, 2)
			getPermutation(3, 3)
			getPermutation(3, 6)
		}
	})
}
