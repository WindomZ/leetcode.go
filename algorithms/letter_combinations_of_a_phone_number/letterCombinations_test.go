package lettercombinationsofaphonenumber

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{"0"}, letterCombinations("0"))
	assert.Equal(t, []string{"01"}, letterCombinations("01"))
	assert.Equal(t, []string{
		"01a", "01b", "01c"},
		letterCombinations("012"),
	)
	assert.Equal(t, []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		letterCombinations("23"),
	)
	assert.Equal(t, []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		letterCombinations("23x4"),
	)
}

func Benchmark_letterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("")
		letterCombinations("01")
		letterCombinations("24")
		letterCombinations("159")
	}
}
