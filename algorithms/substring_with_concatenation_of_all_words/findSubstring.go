package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return []int{}
	}

	lenWs, lenW, end := len(words), len(words[0]), len(s)-len(words)*len(words[0])
	remain := make(map[string]int, lenWs)
	result := make([]int, 0, end/lenW)

	for offset := 0; offset < lenW; offset++ { // offset = [0, length of word)
		// reset remain mapping
		for _, word := range words {
			remain[word] = 0
		}
		for _, word := range words {
			remain[word]++
		}
		// start: the starting indices
		// count: the count of matching words
		for start, count := offset, 0; start <= end; {
			word := s[start+count*lenW : start+(count+1)*lenW] // the word
			if n, ok := remain[word]; ok {                     // if match word
				if n == 0 { // if match word more than once
					remain[s[start:start+lenW]]++ // restore the first word count
					start += lenW                 // move start to next word
					count--                       // restore one count of matching words
				} else { // match word exactly once
					remain[word]--
					if count++; count == lenWs {
						result = append(result, start) // append the one of indices
						remain[s[start:start+lenW]]++  // restore the first word count
						start += lenW                  // move start to next word
						count--                        // restore one count of matching words
					}
				}
			} else { // has any intervening characters
				if count != 0 {
					// reset remain mapping
					for _, word := range words {
						remain[word] = 0
					}
					for _, word := range words {
						remain[word]++
					}
				}
				start += lenW * (count + 1) // reset start
				count = 0                   // reset count
			}
		}
	}

	return result
}
