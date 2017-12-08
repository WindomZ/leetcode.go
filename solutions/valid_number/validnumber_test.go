package validnumber

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isNumber(t *testing.T) {
	assert.Equal(t, false, isNumber(""))
	assert.Equal(t, true, isNumber("0"))
	assert.Equal(t, true, isNumber("1"))
	assert.Equal(t, true, isNumber(" 1"))
	assert.Equal(t, true, isNumber("1 "))
	assert.Equal(t, true, isNumber(" 1 "))
	assert.Equal(t, true, isNumber(" 0.1"))
	assert.Equal(t, true, isNumber("0.1 "))
	assert.Equal(t, true, isNumber(" 0.1 "))
	assert.Equal(t, true, isNumber("1.2"))
	assert.Equal(t, false, isNumber("abc"))
	assert.Equal(t, false, isNumber(" abc"))
	assert.Equal(t, false, isNumber("abc "))
	assert.Equal(t, false, isNumber(" abc "))
	assert.Equal(t, false, isNumber("1 a"))
	assert.Equal(t, false, isNumber(" 1 a"))
	assert.Equal(t, false, isNumber("1 a "))
	assert.Equal(t, false, isNumber(" 1 a "))
	assert.Equal(t, true, isNumber("2e10"))
	assert.Equal(t, true, isNumber(" 2e10"))
	assert.Equal(t, true, isNumber("2e10 "))
	assert.Equal(t, true, isNumber(" 2e10 "))
	assert.Equal(t, false, isNumber(" 2e 10 "))
	assert.Equal(t, false, isNumber("2e 10"))
	assert.Equal(t, false, isNumber("2e1 0"))
	assert.Equal(t, true, isNumber("1.2e3"))
	assert.Equal(t, true, isNumber(" 1.2e3"))
	assert.Equal(t, true, isNumber("1.2e3 "))
	assert.Equal(t, true, isNumber(" 1.2e3 "))
	assert.Equal(t, false, isNumber(" 1.2 e3 "))
	assert.Equal(t, true, isNumber("+1.2e3"))
	assert.Equal(t, true, isNumber(" +1.2e3"))
	assert.Equal(t, true, isNumber("+1.2e3 "))
	assert.Equal(t, true, isNumber(" +1.2e3 "))
}

func Benchmark_isNumber(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isNumber("0")
			isNumber("1.2")
			isNumber("+1.2e3")
			isNumber(" 0.1 ")
			isNumber(" abc ")
			isNumber(" 2e10 ")
			isNumber(" 2e 10 ")
			isNumber(" +1.2e3 ")
		}
	})
}
