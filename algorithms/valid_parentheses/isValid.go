package valid_parentheses

func isValid(s string) bool {
	if l := len(s); l >= 1 {
		ri := 0
		r := make([]byte, l)
		for si := 0; si < l; si++ {
			switch s[si] {
			case '(':
				r[ri] = ')'
				ri++
			case '[':
				r[ri] = ']'
				ri++
			case '{':
				r[ri] = '}'
				ri++
			default:
				if ri--; ri < 0 || r[ri] != s[si] {
					return false
				}
			}
		}
		return ri == 0
	}
	return true
}
