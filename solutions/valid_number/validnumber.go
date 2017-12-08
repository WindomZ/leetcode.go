package validnumber

// define a DFA
// 'd' => "digit"
// 'b' => "blank"
// 's' => "sign"
// 'e' => "e"
// '.' => "."
var states = []map[int32]int{
	{}, // 0
	{'b': 1, 's': 2, 'd': 3, '.': 4}, // 1
	{'d': 3, '.': 4},                 // 2
	{'d': 3, '.': 5, 'e': 6, 'b': 9}, // 3
	{'d': 5},                 // 4
	{'d': 5, 'e': 6, 'b': 9}, // 5
	{'s': 7, 'd': 8},         // 6
	{'d': 8},                 // 7
	{'d': 8, 'b': 9},         // 8
	{'b': 9},                 // 9
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
