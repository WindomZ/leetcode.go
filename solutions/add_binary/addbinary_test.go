package addbinary

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, "0", addBinary("0", "0"))
	assert.Equal(t, "1", addBinary("0", "1"))
	assert.Equal(t, "10", addBinary("1", "1"))
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "110", addBinary("11", "11"))
}

func Benchmark_addBinary(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addBinary("0", "0")
			addBinary("0", "1")
			addBinary("1", "1")
			addBinary("11", "1")
			addBinary("11", "11")
		}
	})
}
