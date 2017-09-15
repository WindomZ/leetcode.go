package implementstrstr

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_strStr(t *testing.T) {
	assert.Equal(t, 0, strStr("", ""))
	assert.Equal(t, 0, strStr("A", ""))
	assert.Equal(t, 1, strStr("ABC", "B"))
	assert.Equal(t, 2, strStr("ABCDE", "CD"))
	assert.Equal(t, 3, strStr("ABCDEFG", "DEF"))
	assert.Equal(t, 8, strStr("ABCDEFGHI", "I"))
	assert.Equal(t, -1, strStr("ABCDEFGHIJK", "HIK"))
}

func Benchmark_strStr(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			strStr("", "")
			strStr("A", "")
			strStr("ABC", "B")
			strStr("ABCDE", "CD")
			strStr("ABCDEFG", "DEF")
			strStr("ABCDEFGHI", "I")
			strStr("ABCDEFGHIJK", "HIK")
		}
	})
}
