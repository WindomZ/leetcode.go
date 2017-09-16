package generateparentheses

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
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generateParenthesis(0)
			generateParenthesis(1)
			generateParenthesis(2)
			generateParenthesis(3)
		}
	})
}
