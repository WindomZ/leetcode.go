package searchforarange

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
	} else if len(nums) >= 2 {
		low, high := 0, len(nums)-1
		// Search for the left one using binary search.
		for low < high {
			middle := (low + high) / 2
			if nums[middle] < target {
				low = middle + 1
			} else {
				high = middle
			}
		}
		if nums[low] == target {
			result[0] = low
			high = len(nums) - 1
			// Search for the right one using binary search.
			for low < high {
				middle := (low+high)/2 + 1 // middle value trends to the right
				if nums[middle] > target {
					high = middle - 1
				} else {
					low = middle
				}
			}
			result[1] = high
		}
	}
	return result
}
