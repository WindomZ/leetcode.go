package multiplystrings

func multiply(num1 string, num2 string) string {
	sum := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		var carry, bit byte
		// multiply
		for j := len(num2) - 1; j >= 0; j-- {
			bit = byte(num1[i]-'0')*byte(num2[j]-'0') + carry // bit=num1[i]*num2[j]+carry
			if sum[i+j+1] != 0 {
				bit += byte(sum[i+j+1] - '0') // supplements last i+j+1 bit number
			}
			sum[i+j+1] = bit%10 + '0' // compute new i+j+1 bit number
			carry = bit / 10          // compute next carry bit
		}
		// carry bit
		if carry != 0 {
			sum[i] = carry + '0'
		}
	}

	// get the legal highest bit
	for i := 0; i < len(sum); i++ {
		if sum[i] != 0 && sum[i]-'0' > 0 {
			return string(sum[i:])
		}
	}

	return "0"
}
