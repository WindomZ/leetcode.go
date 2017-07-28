package regular_expression_matching

func isMatch(s string, p string) bool {
	if s != p {
		pl := len(p)
		if pl == 0 || p[0] == '*' {
			return false
		}
		sl := len(s)
		if sl == 0 {
			if pl%2 != 0 {
				return false
			}
			for i := 1; i < pl; i += 2 {
				if p[i] != '*' {
					return false
				}
			}
			return true
		}

		// init matrix r
		r := make([][]bool, sl+1)
		for si := 0; si <= sl; si++ {
			r[si] = make([]bool, pl+1)
		}
		r[0][0] = true

		for pi := 2; pi <= pl; pi++ {
			r[0][pi] = p[pi-1] == '*' && r[0][pi-2]
		}

		// calc matrix r
		for si := 1; si <= sl; si++ {
			for pi := 1; pi <= pl; pi++ {
				if p[pi-1] == '*' { // p_byte == '*'
					// r[si][pi-2] is true, p['a*'] matches zero.
					// r[si-1][pi] is true, last s and p['?*'] matches.
					// s[si-1] == p[pi-2] is true, s['a'] and p['a*'] matches one/more.
					// p[pi-2] == '.' is true, s['a'] and p['.*'] matches one/more.
					r[si][pi] = r[si][pi-2] || (r[si-1][pi] && (s[si-1] == p[pi-2] || p[pi-2] == '.'))
				} else {
					// r[si-1][pi-1] is true, last s and last p matches.
					// s[si-1] == p[pi-1] is true, s['a'] and p['a'] matches.
					// p[pi-1] == '.' is true, s['a'] and p['.'] matches.
					r[si][pi] = r[si-1][pi-1] && (s[si-1] == p[pi-1] || p[pi-1] == '.')
				}
			}
		}

		return r[sl][pl]
	}
	return true
}
