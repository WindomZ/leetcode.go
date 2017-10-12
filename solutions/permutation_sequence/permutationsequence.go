package permutationsequence

func getPermutation(n int, k int) string {
	if k <= 0 {
		return ""
	}
	s := make([]byte, n)
	i, j, f := 1, 0, 1
	for ; i <= n; i++ {
		f *= i
		s[i-1] += byte('0' + i) // make s become 1234...n
	}
	for i, k = 0, k-1; i < n; i++ {
		f /= n - i
		j = i + k/f // calculate and put at s[i]
		temp := s[j]
		// remove c by shifting to cover up (adjust the right part).
		for ; j > i; j-- {
			s[j] = s[j-1]
		}
		k %= f
		s[i] = temp
	}
	return string(s)
}
