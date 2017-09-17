package longestcommonprefix

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{}))
	assert.Equal(t, "", longestCommonPrefix([]string{"ca", "a"}))
	assert.Equal(t, "a", longestCommonPrefix([]string{"aa", "a"}))
	assert.Equal(t, "a", longestCommonPrefix([]string{"abc", "acb"}))
	assert.Equal(t, "ab", longestCommonPrefix([]string{"abc", "abd"}))
	assert.Equal(t, "ab", longestCommonPrefix([]string{"abcd", "abce", "abde"}))
	assert.Equal(t, "abc", longestCommonPrefix([]string{"abcd", "abcde", "abcdef", "abcefg"}))
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			longestCommonPrefix([]string{})
			longestCommonPrefix([]string{"", "acb"})
			longestCommonPrefix([]string{"abc", "cba"})
			longestCommonPrefix([]string{"abc", "acb"})
			longestCommonPrefix([]string{"abc", "abd"})
			longestCommonPrefix([]string{"abcd", "abce", "abde"})
			longestCommonPrefix([]string{"abcd", "abce", "bdea", "abde"})
			longestCommonPrefix([]string{"abcd", "abcde", "abcdef", "abcefg"})
		}
	})
}
