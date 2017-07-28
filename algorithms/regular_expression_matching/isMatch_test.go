package regular_expression_matching

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isMatch(t *testing.T) {
	assert.Equal(t, false, isMatch("a", ""))
	assert.Equal(t, false, isMatch("", "a"))
	assert.Equal(t, false, isMatch("a", "*"))
	assert.Equal(t, false, isMatch("a", "ab*a"))
	assert.Equal(t, false, isMatch("ab", "aa"))
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, false, isMatch("aa", "a.b"))
	assert.Equal(t, false, isMatch("aaa", "aa"))
	assert.Equal(t, false, isMatch("acb", "a*b"))
	assert.Equal(t, false, isMatch("aab", "b.*"))
	assert.Equal(t, false, isMatch("aacc", "a.*b"))

	assert.Equal(t, true, isMatch("", ""))
	assert.Equal(t, true, isMatch("", ".*"))
	assert.Equal(t, true, isMatch("", "c*"))
	assert.Equal(t, true, isMatch("", "a*b*c*"))
	assert.Equal(t, true, isMatch("aa", "aa"))
	assert.Equal(t, true, isMatch("aa", "a."))
	assert.Equal(t, true, isMatch("aa", "a*"))
	assert.Equal(t, true, isMatch("aa", ".*"))
	assert.Equal(t, true, isMatch("ab", ".*"))
	assert.Equal(t, true, isMatch("aaa", "a*a"))
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
	assert.Equal(t, true, isMatch("aacb", "a.*b"))
	assert.Equal(t, true, isMatch("aaa", "ab*a*c*a"))
	assert.Equal(t, true, isMatch("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s"))
	assert.Equal(t, true, isMatch("abcdede", "ab.*de"))
}

func Benchmark_isMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatch("", "")
		isMatch("aa", "aa")
		isMatch("aa", "a.")
		isMatch("ab", ".*")
		isMatch("aab", "c*a*b")
		isMatch("acdefgb", "a.*b")
	}
}
