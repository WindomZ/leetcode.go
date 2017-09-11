package wildcardmatching

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isMatch(t *testing.T) {
	assert.Equal(t, true, isMatch("", ""))
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "aa"))
	assert.Equal(t, false, isMatch("aaa", "aa"))
	assert.Equal(t, true, isMatch("aa", "*"))
	assert.Equal(t, true, isMatch("aa", "a*"))
	assert.Equal(t, true, isMatch("ab", "?*"))
	assert.Equal(t, false, isMatch("aab", "c*a*b"))
	assert.Equal(t, true, isMatch("aab", "*a*b**"))
}

func Benchmark_isMatch(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isMatch("", "")
			isMatch("aa", "a")
			isMatch("aa", "aa")
			isMatch("aaa", "aa")
			isMatch("aa", "*")
			isMatch("aa", "a*")
			isMatch("ab", "?*")
			isMatch("aab", "c*a*b")
			isMatch("aab", "*a*b**")
		}
	})
}
