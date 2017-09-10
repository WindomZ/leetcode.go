package longestvalidparentheses

func longestValidParentheses(s string) int {
	l := 0
	if len(s) >= 1 {
		i, stack := 0, make([]int, len(s)+1)
		stack[0] = -1 // first index
		for si := 0; si < len(s); si++ {
			if s[si] == '(' { // if '('
				i++ // push
				stack[i] = si
			} else { // if ')'
				i--        // pop
				if i < 0 { // keep first index
					i = 0 // push
					stack[i] = si
				} else if tmp := si - stack[i]; tmp > l {
					l = tmp // max length
				}
			}
		}
	}
	return l
}
