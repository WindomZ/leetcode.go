package validnumber

// define a DFA
// 'd' => "digit"
// 'b' => "blank"
// 's' => "sign"
// 'e' => "e"
// '.' => "."
// from 0 to 9
var states = []map[int32]int{
	{},
	{'b': 1, 's': 2, 'd': 3, '.': 4},
	{'d': 3, '.': 4},
	{'d': 3, '.': 5, 'e': 6, 'b': 9},
	{'d': 5},
	{'d': 5, 'e': 6, 'b': 9},
	{'s': 7, 'd': 8},
	{'d': 8},
	{'d': 8, 'b': 9},
	{'b': 9},
}

// isNumber @see valid_number_dfa.svg
func isNumber(s string) bool {
	index := 1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			c = 'd'
		} else if c == ' ' {
			c = 'b'
		} else if c == '+' || c == '-' {
			c = 's'
		}
		if v, ok := states[index][c]; ok {
			index = v
		} else {
			return false
		}
	}
	// valid states are end on digit, after dot, digit after e, or white space after valid input
	return index == 3 || index == 5 || index == 8 || index == 9
}
