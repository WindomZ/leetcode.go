package letter_combinations_of_a_phone_number

var charsmap [10]string = [10]string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	r := []string{""}
	for i := 0; i < len(digits); i++ {
		idx := digits[i] - '0'
		if idx < 0 || idx >= 10 {
			break
		}

		tmp := r
		chars := charsmap[idx]
		r = make([]string, len(r)*len(chars))
		idx = 0
		for _, tv := range tmp {
			for _, cv := range chars {
				r[idx] = tv + string(cv)
				idx++
			}
		}
	}
	return r
}
