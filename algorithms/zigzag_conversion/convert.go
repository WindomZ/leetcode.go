package zigzagconversion

func convert(s string, numRows int) string {
	if l := len(s); numRows >= 2 && numRows < l {
		var idx int
		data := make([]byte, l)
		interval := 2*numRows - 2

		// first row
		for i := 0; i < l; i += interval {
			data[idx] = s[i]
			idx++
		}

		// middle rows
		for i := 1; i <= numRows-2; i++ {
			data[idx] = s[i]
			idx++
			for j := interval; j-i < l; j += interval {
				data[idx] = s[j-i]
				idx++
				if j+i < l {
					data[idx] = s[j+i]
					idx++
				}
			}
		}

		// last row
		for i := numRows - 1; i < l; i += interval {
			data[idx] = s[i]
			idx++
		}

		return string(data)
	}
	return s
}
