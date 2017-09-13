package groupanagrams

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{}, groupAnagrams([]string{}))

	r := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert.Contains(t, r, []string{"eat", "tea", "ate"})
	assert.Contains(t, r, []string{"tan", "nat"})
	assert.Contains(t, r, []string{"bat"})

	r = groupAnagrams([]string{"are", "bat", "ear", "code", "tab", "era"})
	assert.Contains(t, r, []string{"are", "ear", "era"})
	assert.Contains(t, r, []string{"bat", "tab"})
	assert.Contains(t, r, []string{"code"})
}

func Benchmark_groupAnagrams(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			groupAnagrams([]string{})
			groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
			groupAnagrams([]string{"are", "bat", "ear", "code", "tab", "era"})
		}
	})
}
