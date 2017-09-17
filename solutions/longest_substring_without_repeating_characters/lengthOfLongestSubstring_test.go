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
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lengthOfLongestSubstring("abcabcbb")
			lengthOfLongestSubstring("bbbbb")
			lengthOfLongestSubstring("pwwkew")
		}
	})
}
