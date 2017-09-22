package lengthoflastword

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	l, tail := 0, len(s)-1
	for ; tail >= 0 && s[tail] == ' '; tail-- { // exclude tail spaces
	}
	for ; tail >= 0 && s[tail] != ' '; tail-- { // the length of last word in the string
		l++
	}
	return l
}
