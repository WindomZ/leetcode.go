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
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			longestValidParentheses("")
			longestValidParentheses("()")
			longestValidParentheses("()()")
			longestValidParentheses("(()()")
			longestValidParentheses("()(()")
			longestValidParentheses(")()())")
			longestValidParentheses("(()(((()")
			longestValidParentheses(")()(((())))(")
		}
	})
}
