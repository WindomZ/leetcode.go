package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if n := len(s); n >= 1 {
		var l, tmp, index, offset int
		var offsets [256]int
		for ; index < n; index++ {
			if tmp = offsets[s[index]]; tmp >= offset {
				offset = tmp // no repeating character offset
			}
			if tmp = index - offset + 1; tmp > l {
				l = tmp // no repeating character length
			}
			offsets[s[index]] = index + 1 // record the character offset
		}
		return l
	}
	return 0
}
