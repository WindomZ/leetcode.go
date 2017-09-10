package validparentheses

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isValid(t *testing.T) {
	assert.Equal(t, false, isValid("("))
	assert.Equal(t, false, isValid(")"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("[)"))
	assert.Equal(t, false, isValid("([}"))
	assert.Equal(t, false, isValid("([)]"))

	assert.Equal(t, true, isValid(""))
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("[]"))
	assert.Equal(t, true, isValid("{}"))
	assert.Equal(t, true, isValid("{[()]}"))
	assert.Equal(t, true, isValid("()[]{}"))
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("")
		isValid("([)]")
		isValid("{[()]}")
		isValid("()[]{}")
		isValid("()[]{}[]()")
		isValid("()[]{[()]}[]()")
	}
}
