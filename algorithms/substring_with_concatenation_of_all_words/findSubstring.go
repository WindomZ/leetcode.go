package substring_with_concatenation_of_all_words

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return []int{}
	}

	len_ws, len_w, end := len(words), len(words[0]), len(s)-len(words)*len(words[0])
	remain := make(map[string]int, len_ws)
	result := make([]int, 0, end/len_w)

	for offset := 0; offset < len_w; offset++ { // offset = [0, length of word)
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
			word := s[start+count*len_w : start+(count+1)*len_w] // the word
			if n, ok := remain[word]; ok {                       // if match word
				if n == 0 { // if match word more than once
					remain[s[start:start+len_w]]++ // restore the first word count
					start += len_w                 // move start to next word
					count--                        // restore one count of matching words
				} else { // match word exactly once
					remain[word]--
					if count++; count == len_ws {
						result = append(result, start) // append the one of indices
						remain[s[start:start+len_w]]++ // restore the first word count
						start += len_w                 // move start to next word
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
				start += len_w * (count + 1) // reset start
				count = 0                    // reset count
			}
		}
	}

	return result
}
