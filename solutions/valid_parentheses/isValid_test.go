package validparentheses

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isValid(t *testing.T) {
	assert.Equal(t, false, isValid("("))
	assert.Equal(t, false, isValid(")"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("[)"))
	assert.Equal(t, false, isValid("([}"))
	assert.Equal(t, false, isValid("([)]"))

	assert.Equal(t, true, isValid(""))
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("[]"))
	assert.Equal(t, true, isValid("{}"))
	assert.Equal(t, true, isValid("{[()]}"))
	assert.Equal(t, true, isValid("()[]{}"))
}

func Benchmark_isValid(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isValid("")
			isValid("([)]")
			isValid("{[()]}")
			isValid("()[]{}")
			isValid("()[]{}[]()")
			isValid("()[]{[()]}[]()")
		}
	})
}
