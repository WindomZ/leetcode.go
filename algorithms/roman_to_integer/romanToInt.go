package roman_to_integer

func romanToInt(s string) int {
	if l := len(s); l >= 1 {
		var last, current, result int
		for i := 0; i < l; i++ {
			switch s[i] {
			case 'M':
				result += 1000
				current = 7
			case 'D':
				result += 500
				current = 6
			case 'C':
				result += 100
				current = 5
			case 'L':
				result += 50
				current = 4
			case 'X':
				result += 10
				current = 3
			case 'V':
				result += 5
				current = 2
			default: // 'I'
				result += 1
				current = 1
			}
			if last != 0 && last < current {
				switch last*10 + current {
				case 12, 13: // "IV"/"IX"
					result -= 2
				case 34, 35: // "XL"/"XC"
					result -= 20
				case 56, 57: // "CD"/"CM"
					result -= 200
				}
			}
			last = current
		}
		return result
	}
	return 0
}
