package decodeways

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_numDecodings(t *testing.T) {
	assert.Equal(t, 0, numDecodings(""))
	assert.Equal(t, 1, numDecodings("1"))
	assert.Equal(t, 1, numDecodings("2"))
	assert.Equal(t, 2, numDecodings("11"))
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 2, numDecodings("26"))
	assert.Equal(t, 1, numDecodings("27"))
	assert.Equal(t, 0, numDecodings("100"))
	assert.Equal(t, 1, numDecodings("101"))
	assert.Equal(t, 3, numDecodings("112"))
	assert.Equal(t, 3, numDecodings("126"))
	assert.Equal(t, 2, numDecodings("127"))
	assert.Equal(t, 1, numDecodings("1010"))
	assert.Equal(t, 4, numDecodings("13142"))
	assert.Equal(t, 3, numDecodings("123456"))
}

func Benchmark_numDecodings(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			numDecodings("1")
			numDecodings("12")
			numDecodings("27")
			numDecodings("100")
			numDecodings("126")
			numDecodings("1010")
			numDecodings("13142")
			numDecodings("123456")
		}
	})
}
