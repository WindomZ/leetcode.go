package integer_to_roman

var (
	M = [4]string{"", "M", "MM", "MMM"}
	C = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	if num >= 1 && num <= 3999 {
		if num >= 1000 {
			return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
		} else if num >= 100 {
			return C[num/100] + X[(num%100)/10] + I[num%10]
		}
		return X[(num%100)/10] + I[num%10]
	}
	return ""
}
