package integer_to_roman

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_intToRoman(t *testing.T) {
	assert.Equal(t, "", intToRoman(0))
	assert.Equal(t, "", intToRoman(4000))

	assert.Equal(t, "I", intToRoman(1))
	assert.Equal(t, "X", intToRoman(10))
	assert.Equal(t, "C", intToRoman(100))
	assert.Equal(t, "M", intToRoman(1000))
	assert.Equal(t, "MCXI", intToRoman(1111))
	assert.Equal(t, "MCX", intToRoman(1110))
	assert.Equal(t, "MMMCMXCIX", intToRoman(3999))
}

func Benchmark_intToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(0)
		intToRoman(9)
		intToRoman(10)
		intToRoman(99)
		intToRoman(100)
		intToRoman(999)
		intToRoman(1000)
		intToRoman(3999)
	}
}
