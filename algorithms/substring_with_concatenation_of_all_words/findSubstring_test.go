package substring_with_concatenation_of_all_words

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_findSubstring(t *testing.T) {
	assert.Equal(t, []int{}, findSubstring("", []string{""}))
	assert.Equal(t, []int{0}, findSubstring("a", []string{"a"}))
	assert.Equal(t, []int{1}, findSubstring("acaacc",
		[]string{"ca", "ac"}))
	assert.Equal(t, []int{1}, findSubstring("ababaab",
		[]string{"ab", "ba", "ba"}))
	assert.Equal(t, []int{1, 3}, findSubstring("abaababbaba",
		[]string{"ab", "ba", "ab", "ba"}))
	assert.Equal(t, []int{0, 2, 1}, findSubstring("aaaaaaaa",
		[]string{"aa", "aa", "aa"}))
	assert.Equal(t, []int{0, 9}, findSubstring("barfoothefoobarman",
		[]string{"foo", "bar"}))
	assert.Equal(t, []int{8}, findSubstring("wordgoodgoodgoodbestword",
		[]string{"word", "good", "best", "good"}))
	assert.Equal(t, []int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman",
		[]string{"bar", "foo", "the"}))
}

func Benchmark_findSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findSubstring("", []string{""})
		findSubstring("acaacc",
			[]string{"ca", "ac"})
		findSubstring("ababaab",
			[]string{"ab", "ba", "ba"})
	}
}
