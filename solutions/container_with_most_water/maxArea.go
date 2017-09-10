package containerwithmostwater

func maxArea(height []int) int {
	if hs := len(height); hs >= 2 {
		var max, h, l, r int = 0, 0, 0, hs - 1
		for l < r {
			// get the min height
			if height[l] >= height[r] {
				h = height[r]
				r--
			} else {
				h = height[l]
				l++
			}
			// get the max area
			if h *= r - l + 1; h > max {
				max = h
			}
		}
		return max
	}
	return 0
}
