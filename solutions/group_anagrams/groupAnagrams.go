package groupanagrams

// primes assign a prime number for a to z
var primes = [27]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 107}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	var k int64
	m := make(map[int64][]string, len(strs))
	for _, s := range strs {
		k = 1
		for _, c := range s {
			// multiplies all prime numbers together to form a hash value
			k *= primes[c-'a']
		}
		// appends to a map with k key
		if _, ok := m[k]; !ok {
			m[k] = make([]string, 0, 3)
		}
		m[k] = append(m[k], s)
	}

	// generates each group of Strings with the unique k key
	k = 0
	r := make([][]string, len(m))
	for _, v := range m {
		r[k] = v
		k++
	}
	return r
}
