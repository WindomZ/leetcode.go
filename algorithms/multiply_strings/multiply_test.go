package multiplystrings

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_multiply(t *testing.T) {
	assert.Equal(t, "0", multiply("", ""))
	assert.Equal(t, "0", multiply("1", "0"))
	assert.Equal(t, "1", multiply("1", "1"))
	assert.Equal(t, "2", multiply("1", "2"))
	assert.Equal(t, "20", multiply("10", "2"))
	assert.Equal(t, "276", multiply("12", "23"))
	assert.Equal(t, "5535", multiply("123", "45"))
	assert.Equal(t, "98901", multiply("999", "99"))
}

func Benchmark_multiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiply("1", "0")
		multiply("1", "1")
		multiply("10", "2")
		multiply("12", "23")
		multiply("123", "45")
		multiply("999", "99")
	}
}
