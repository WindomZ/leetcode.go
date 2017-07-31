package generate_parentheses

func generateParenthesis(n int) []string {
	if n >= 1 {
		var r []string
		generate(&r, "(", n-1, 1)
		return r
	}
	return []string{}
}

func generate(r *[]string, s string, open, close int) {
	if open == 0 && close == 0 {
		*r = append(*r, s)
		return
	}
	if open > 0 {
		generate(r, s+"(", open-1, close+1)
	}
	if close > 0 {
		generate(r, s+")", open, close-1)
	}
}
