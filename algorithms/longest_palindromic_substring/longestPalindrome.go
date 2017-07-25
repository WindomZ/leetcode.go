package longest_palindromic_substring

func longestPalindrome(s string) string {
	if ls := len(s) - 1; ls >= 0 {
		var l, r, t int
		for i := 0; i < ls; i++ {
			odd := longPalindrome(s, i, i)
			even := longPalindrome(s, i, i+1)
			if even <= odd {
				if odd != 1 && odd > t {
					l = i - (odd-1)/2
					r = i + odd/2
					t = r - l + 1
				}
			} else if even > t {
				l = i - even/2 + 1
				r = i + even/2
				t = r - l + 1
			}
		}
		return s[l : r+1]
	}
	return s
}

func longPalindrome(s string, l, r int) int {
	size := len(s)
	for l >= 0 && r < size {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return r - l - 1
}
