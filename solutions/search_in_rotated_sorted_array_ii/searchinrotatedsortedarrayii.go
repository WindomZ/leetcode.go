package searchinrotatedsortedarrayii

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for m := 0; l < r; {
		m = (l + r) / 2
		if nums[m] == target {
			return true
		}
		if nums[m] > nums[r] {
			if nums[m] > target && nums[l] <= target {
				r = m
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] {
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m
			}
		} else {
			r--
		}
	}
	return nums[l] == target
}
