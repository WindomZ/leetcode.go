package addbinary

func addBinary(a string, b string) string {
	var s string

	var c uint8
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || c == 1; {
		if i >= 0 {
			c += a[i] - '0' // sum a binary string
			i--
		}
		if j >= 0 {
			c += b[j] - '0' // sum b binary string
			j--
		}
		s = string(c%2+'0') + s // count s binary string
		c /= 2                  // carry 1
	}

	return s
}
