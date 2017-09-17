package romantointeger

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	assert.Equal(t, 0, romanToInt(""))
	assert.Equal(t, 1, romanToInt("I"))
	assert.Equal(t, 10, romanToInt("X"))
	assert.Equal(t, 100, romanToInt("C"))
	assert.Equal(t, 1000, romanToInt("M"))
	assert.Equal(t, 1110, romanToInt("MCX"))
	assert.Equal(t, 1111, romanToInt("MCXI"))
	assert.Equal(t, 2111, romanToInt("MMCXI"))
	assert.Equal(t, 3011, romanToInt("MMCMCXI"))
	assert.Equal(t, 3999, romanToInt("MMMCMXCIX"))
	assert.Equal(t, 2443, romanToInt("MCMCDXCXLIXIV"))
}

func Benchmark_romanToInt(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			romanToInt("")
			romanToInt("I")
			romanToInt("MCX")
			romanToInt("MCXI")
			romanToInt("MMCXI")
			romanToInt("MMCMCXI")
			romanToInt("MMMCMXCIX")
			romanToInt("MCMCDXCXLIXIV")
		}
	})
}
