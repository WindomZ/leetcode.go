package palindrome_number

func isPalindrome(x int) bool {
	if x <= -1 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	y := 0
	for ; x > y; x = x / 10 {
		y = y*10 + x%10
	}

	return x == y || x == y/10
}
