package decodeways

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1 // have one way to decode
	if s[len(s)-1] != '0' {
		dp[len(s)-1] = 1 //  no digit start with '0'
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' { // no digit start with '0'
			continue
		} else {
			if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
				dp[i] = dp[i+1] + dp[i+2] // check two digit combination: s[i:2+2] <= 26
			} else {
				dp[i] = dp[i+1] // check one digit combination
			}
		}
	}
	return dp[0]
}
