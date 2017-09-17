package zigzagconversion

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_convert(t *testing.T) {
	assert.Equal(t, "", convert("", 1))
	assert.Equal(t, "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
	assert.Equal(t, "PYAIHRNAPLSIIG", convert("PAYPALISHIRING", 2))
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))

	assert.Equal(t, "ABCED", convert("ABCDE", 4))
}

func Benchmark_convert(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			convert("PAYPALISHIRING", 1)
			convert("PAYPALISHIRING", 2)
			convert("PAYPALISHIRING", 3)
			convert("PAYPALISHIRING", 4)
			convert("PAYPALISHIRING", 5)
		}
	})
}
