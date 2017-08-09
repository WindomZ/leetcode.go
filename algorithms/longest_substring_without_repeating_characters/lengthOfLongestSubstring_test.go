package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
		lengthOfLongestSubstring("bbbbb")
		lengthOfLongestSubstring("pwwkew")
	}
}
