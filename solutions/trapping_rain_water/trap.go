package trappingrainwater

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var result int
	for left, right, leftMax, rightMax := 0, len(height)-1, 0, 0; left < right; {
		// from low to high
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left] // find the left max height
			} else {
				result += leftMax - height[left] // compute left drop water
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right] // find the right max height
			} else {
				result += rightMax - height[right] // compute right drop water
			}
			right--
		}
	}
	return result
}
