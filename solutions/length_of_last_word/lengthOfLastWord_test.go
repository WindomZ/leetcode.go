package lengthoflastword

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	assert.Equal(t, 0, lengthOfLastWord(""))
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 5, lengthOfLastWord("  Hello World  "))
	assert.Equal(t, 1, lengthOfLastWord("x xx xxx xx x"))
	assert.Equal(t, 3, lengthOfLastWord("xxx XXX !@#"))
}

func Benchmark_lengthOfLastWord(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lengthOfLastWord("")
			lengthOfLastWord("Hello World")
			lengthOfLastWord("  Hello World  ")
			lengthOfLastWord("x xx xxx xx x")
			lengthOfLastWord("xxx XXX !@#")
		}
	})
}
