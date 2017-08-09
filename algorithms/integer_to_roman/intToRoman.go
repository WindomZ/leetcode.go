package integertoroman

var (
	m = [4]string{"", "m", "MM", "MMM"}
	c = [10]string{"", "c", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	x = [10]string{"", "x", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	i = [10]string{"", "i", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	if num >= 1 && num <= 3999 {
		if num >= 1000 {
			return m[num/1000] + c[(num%1000)/100] + x[(num%100)/10] + i[num%10]
		} else if num >= 100 {
			return c[num/100] + x[(num%100)/10] + i[num%10]
		}
		return x[(num%100)/10] + i[num%10]
	}
	return ""
}
