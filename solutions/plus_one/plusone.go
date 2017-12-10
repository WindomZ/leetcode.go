package plusone

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] <= 8 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
