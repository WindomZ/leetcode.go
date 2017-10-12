package permutationsequence

func getPermutation(n int, k int) string {
	if k <= 0 {
		return ""
	}
	s := make([]byte, n)
	i, j, f := 1, 0, 1
	for ; i <= n; i++ {
		f *= i
		s[i-1] += byte('0' + i) // make s become 1...n
	}
	for i, k = 0, k-1; i < n; i++ {
		f /= n - i
		j = i + k/f
		// move s[j] to s[i] and move s[i...j] to s[i+1...j+1].
		temp := s[j]
		for ; j > i; j-- {
			s[j] = s[j-1]
		}
		k %= f
		s[i] = temp
	}
	return string(s)
}
