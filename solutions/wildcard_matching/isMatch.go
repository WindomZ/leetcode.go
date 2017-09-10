package wildcardmatching

func isMatch(s string, p string) bool {
	var si, pi int
	for starS, starP := 0, -1; si < len(s); {
		if pi < len(p) && matchOne(s[si], p[pi]) { // both characters are match
			si++ // advancing s
			pi++ // advancing p
		} else if pi < len(p) && p[pi] == '*' { // * found, only advancing p
			starP = pi // record starP
			starS = si // record starS
			pi++       // advancing p
		} else if starP != -1 { // last p character was *
			pi = starP + 1 // rollback p from starP, and advancing p
			starS++        // advancing starS, meet the matching *
			si = starS     // rollback s from starS
		} else { // current p and last p are not *
			return false // characters does not match
		}
	}

	// check for remaining characters in p
	for pi < len(p) && p[pi] == '*' {
		pi++ // allow * is empty
	}

	return pi == len(p)
}

func matchOne(b1 byte, b2 byte) bool {
	if b2 == '?' || b1 == b2 {
		return true
	}
	return false
}
