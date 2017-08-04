package implement_strstr

func strStr(haystack string, needle string) int {
	if hl, nl := len(haystack), len(needle); nl == 0 {
		return 0
	} else if hl != 0 {
		for i := 0; i <= hl-nl; i++ {
			j := 0
			for ; j < nl && haystack[i+j] == needle[j]; j++ { // match needle in haystack
			}
			if j == nl { // find needle in haystack
				return i
			}
		}
	}
	return -1
}
