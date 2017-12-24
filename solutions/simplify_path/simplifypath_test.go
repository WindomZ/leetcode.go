package simplifypath

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_simplifyPath(t *testing.T) {
	assert.Equal(t, "/", simplifyPath(""))
	assert.Equal(t, "/home", simplifyPath("/home/"))
	assert.Equal(t, "/home/a/b", simplifyPath("/home/a/b"))
	assert.Equal(t, "/b", simplifyPath("/home/../b"))
	assert.Equal(t, "/home/a/b", simplifyPath("../home/a/b"))
	assert.Equal(t, "/b", simplifyPath("../home/../b"))
	assert.Equal(t, "/c", simplifyPath("/a/./b/../../c/"))
}

func Benchmark_simplifyPath(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			simplifyPath("/home/")
			simplifyPath("/home/a/b")
			simplifyPath("/home/../b")
			simplifyPath("../home/a/b")
			simplifyPath("../home/../b")
			simplifyPath("/a/./b/../../c/")
		}
	})
}
