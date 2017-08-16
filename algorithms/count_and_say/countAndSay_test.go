package countandsay

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	assert.Equal(t, "", countAndSay(0))
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "1211", countAndSay(4))
	assert.Equal(t, "312211", countAndSay(6))
	assert.Equal(t, "1113213211", countAndSay(8))
	assert.Equal(t, "13211311123113112211", countAndSay(10))
}

func Benchmark_countAndSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countAndSay(0)
		countAndSay(1)
		countAndSay(4)
		countAndSay(8)
	}
}
