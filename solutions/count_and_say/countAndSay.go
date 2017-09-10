package countandsay

// examples of nth sequence
// 0
// 1
// 11
// 21
// 1211
// 111221
// 312211
// 13112221
// 1113213211
// 31131211131221
// 13211311123113112211
// ...

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := []byte{'1'}
	for count := 0; n >= 2; n-- { // from n to 2
		temp := res[0]
		r := make([]byte, 0, len(res)*2)
		for _, b := range res {
			if b == temp {
				count++
			} else {
				r = append(r, byte('0'+count), temp)
				temp = b
				count = 1
			}
		}
		if count != 0 {
			r = append(r, byte('0'+count), temp)
		}
		count = 0
		res = r
	}
	return string(res)
}
