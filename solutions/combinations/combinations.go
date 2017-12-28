package combinations

func combine(n int, k int) [][]int {
	if k <= 0 {
		return [][]int{}
	}
	rs := make([][]int, 0, k)
	combineNext(&rs, make([]int, 0, k), 1, n, k)
	return rs
}

func combineNext(rs *[][]int, r []int, s, n, k int) {
	if k == 0 {
		t := make([]int, len(r))
		for k, v := range r {
			t[k] = v
		}
		*rs = append(*rs, t)
		return
	}
	for i := s; i <= n; i++ {
		r = append(r, i)
		combineNext(rs, r, i+1, n, k-1)
		r = r[:len(r)-1]
	}
}
