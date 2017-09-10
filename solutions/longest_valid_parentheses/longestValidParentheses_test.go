package longestvalidparentheses

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	assert.Equal(t, 0, longestValidParentheses(""))
	assert.Equal(t, 0, longestValidParentheses("("))
	assert.Equal(t, 0, longestValidParentheses(")"))

	assert.Equal(t, 2, longestValidParentheses("()"))
	assert.Equal(t, 4, longestValidParentheses("()()"))
	assert.Equal(t, 4, longestValidParentheses("(()()"))
	assert.Equal(t, 2, longestValidParentheses("()(()"))
	assert.Equal(t, 4, longestValidParentheses(")()())"))
	assert.Equal(t, 2, longestValidParentheses("(()(((()"))
	assert.Equal(t, 10, longestValidParentheses(")()(((())))("))
}

func Benchmark_longestValidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses("")
		longestValidParentheses("()")
		longestValidParentheses("()()")
		longestValidParentheses("(()()")
		longestValidParentheses("()(()")
		longestValidParentheses(")()())")
		longestValidParentheses("(()(((()")
		longestValidParentheses(")()(((())))(")
	}
}
