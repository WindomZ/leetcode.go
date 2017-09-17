package reverseinteger

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 0, reverse(0))

	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))

	assert.Equal(t, -1, reverse(-10))
	assert.Equal(t, 101, reverse(10100))

	assert.Equal(t, 0, reverse(1000000003))
	assert.Equal(t, 0, reverse(1534236469))

	assert.Equal(t, 0, reverse(math.MaxInt64))
	assert.Equal(t, 0, reverse(-math.MaxInt64))
	assert.Equal(t, 0, reverse(7085774586302733229))
	assert.Equal(t, 0, reverse(-7085774586302733229))

	assert.Equal(t, 0, reverse(7085774586302733239))
	assert.Equal(t, 0, reverse(-7085774586302733239))
	assert.Equal(t, 0, reverse(1000000000000000039))
}

func Benchmark_reverse(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			reverse(1)
			reverse(12)
			reverse(123)
			reverse(1000000003)
			reverse(1534236469)
		}
	})
}
