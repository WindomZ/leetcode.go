package string_to_integer_atoi

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	assert.Equal(t, 0, myAtoi(""))
	assert.Equal(t, 1, myAtoi("1"))
	assert.Equal(t, 1, myAtoi("+1"))
	assert.Equal(t, -1, myAtoi("-1"))
	assert.Equal(t, 12, myAtoi("+12"))
	assert.Equal(t, -12, myAtoi("-12"))
	assert.Equal(t, 123, myAtoi("+123"))
	assert.Equal(t, -123, myAtoi("-123"))

	assert.Equal(t, 123, myAtoi("  +123 "))
	assert.Equal(t, -123, myAtoi("  -123 "))
	assert.Equal(t, 123, myAtoi("  +123x"))
	assert.Equal(t, -123, myAtoi("  -123x"))

	assert.Equal(t, 2147483647, myAtoi("2147483647"))
	assert.Equal(t, 2147483647, myAtoi("21474836471"))
	assert.Equal(t, -2147483647, myAtoi("-2147483647"))
	assert.Equal(t, -2147483648, myAtoi("-21474836471"))

	assert.Equal(t, -2147483648, myAtoi("-2147483648"))

	assert.Equal(t, -1010023630, myAtoi(" -1010023630o4"))
}

func Benchmark_myAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("")
		myAtoi("-1")
		myAtoi("  123 ")
		myAtoi("+2147483648")
		myAtoi("-2147483648")
	}
}
