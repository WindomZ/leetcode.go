package firstmissingpositive

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// find positive integer, and swap it at the same index value.
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		// find the first missing positive integer.
		if nums[i] != i+1 {
			return i + 1 // index value -> positive integer
		}
	}

	return len(nums) + 1 // index value -> positive integer
}
