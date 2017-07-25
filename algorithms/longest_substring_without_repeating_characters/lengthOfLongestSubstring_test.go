package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring(""), 0)
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
	}
}
