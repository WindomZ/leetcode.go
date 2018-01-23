package restoreipaddresses

func restoreIPAddresses(s string) []string {
	if len(s) <= 3 {
		return []string{}
	}
	result := make([]string, 0, 2)
	a, b, c := len(s)-9, 0, 0
	if a <= 0 {
		a = 1
	}
	var s1, s2, s3, s4 string
	for ; a <= 3; a++ {
		for b = a + 1; b <= a+3; b++ {
			for c = b + 1; c <= b+3 && c < len(s); c++ {
				if c+3 >= len(s) {
					if s1 = s[0:a]; invalid(s1) {
						continue
					}
					if s2 = s[a:b]; invalid(s2) {
						continue
					}
					if s3 = s[b:c]; invalid(s3) {
						continue
					}
					if s4 = s[c:]; invalid(s4) {
						continue
					}
					result = append(result, s1+"."+s2+"."+s3+"."+s4)
				}
			}
		}
	}
	return result
}

func invalid(s string) bool {
	if s[0] == '0' {
		return len(s) != 1
	}
	if len(s) == 3 {
		if s[0] > '2' {
			return true
		}
		if s[0] == '2' {
			if s[1] > '5' {
				return true
			}
			if s[1] == '5' && s[2] > '5' {
				return true
			}
		}
	}
	return false
}
