package searchinsertposition

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	// find the target using binary search.
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
