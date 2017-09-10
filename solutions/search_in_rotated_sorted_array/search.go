package searchinrotatedsortedarray

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
	} else if len(nums) >= 1 {
		low, high := 0, len(nums)-1
		// find the index of the smallest value using binary search.
		for low < high {
			middle := (low + high) / 2
			if nums[middle] > nums[high] {
				low = middle + 1
			} else {
				high = middle
			}
		}
		// the index of the smallest value and also the offset number of places rotated.
		offset := low
		low, high = 0, len(nums)-1
		// find the index of target number using binary search.
		for low <= high { // ensure to find the target value
			middle := (low + high) / 2
			idx := (middle + offset) % len(nums) // the index of real middle value in the sorted nums
			if nums[idx] == target {
				return idx
			}
			if nums[idx] < target {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
	}
	return -1
}
