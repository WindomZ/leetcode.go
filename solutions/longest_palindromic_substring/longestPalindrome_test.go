package longestpalindromicsubstring

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "", longestPalindrome(""))
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "a", longestPalindrome("abcda"))
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "aba", longestPalindrome("abacdfgdcaba"))
}

func Benchmark_longestPalindrome(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			longestPalindrome("")
			longestPalindrome("a")
			longestPalindrome("abcda")
			longestPalindrome("babad")
			longestPalindrome("cbbd")
			longestPalindrome("abacdfgdcaba")
		}
	})
}
