package generate_parentheses

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	assert.Equal(t, []string{}, generateParenthesis(0))
	assert.Equal(t, []string{"()"}, generateParenthesis(1))
	assert.Equal(t, []string{"(())", "()()"}, generateParenthesis(2))
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
	assert.Equal(t, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
		"(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()",
		"()()(())", "()()()()"}, generateParenthesis(4))
}

func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(0)
		generateParenthesis(1)
		generateParenthesis(2)
		generateParenthesis(3)
	}
}
