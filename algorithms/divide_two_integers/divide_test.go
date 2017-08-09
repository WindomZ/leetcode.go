package dividetwointegers

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_divide(t *testing.T) {
	assert.Equal(t, math.MaxInt32, divide(0, 0))
	assert.Equal(t, math.MaxInt32, divide(-math.MaxInt32, -1))
	assert.Equal(t, 0, divide(0, 1))
	assert.Equal(t, 4, divide(12, 3))
	assert.Equal(t, 5, divide(15, 3))
	assert.Equal(t, 5, divide(16, 3))
	assert.Equal(t, 6, divide(18, 3))
	assert.Equal(t, 6, divide(19, 3))
	assert.Equal(t, 5, divide(20, 4))
	assert.Equal(t, 5, divide(23, 4))
	assert.Equal(t, 6, divide(24, 4))
	assert.Equal(t, -6, divide(27, -4))
}

func Benchmark_divide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divide(0, 0)
		divide(-math.MaxInt32, -1)
		divide(0, 1)
		divide(12, 3)
		divide(15, 3)
		divide(16, 3)
		divide(18, 3)
		divide(19, 3)
		divide(20, 4)
		divide(23, 4)
		divide(24, 4)
		divide(27, -4)
	}
}
