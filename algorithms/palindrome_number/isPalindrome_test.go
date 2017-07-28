package palindrome_number

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.Equal(t, false, isPalindrome(-1))
	assert.Equal(t, true, isPalindrome(0))
	assert.Equal(t, true, isPalindrome(1))
	assert.Equal(t, true, isPalindrome(9))
	assert.Equal(t, false, isPalindrome(10))
	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, false, isPalindrome(1210))
	assert.Equal(t, true, isPalindrome(123454321))
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(-1)
		isPalindrome(5)
		isPalindrome(10)
		isPalindrome(121)
		isPalindrome(1210)
		isPalindrome(123454321)
		isPalindrome(123456764321)
	}
}
