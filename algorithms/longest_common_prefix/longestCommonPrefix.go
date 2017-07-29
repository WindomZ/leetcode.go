package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) >= 1 {
		p := strs[0]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < len(p) {
				p = p[:len(strs[i])] // get prefix if longer than strs[i]
			}
			for j := 0; j < len(p); j++ {
				if p[j] != strs[i][j] {
					p = p[:j] // get prefix if is different from strs[i] prefix
					break
				}
			}
		}
		return p
	}
	return ""
}
