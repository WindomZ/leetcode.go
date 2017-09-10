package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums) >= 2 {
		length := 1 // at least one difference
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				nums[length] = nums[i] // remove the duplicates
				length++               // count from different values
			}
		}
		return length
	}
	return len(nums)
}
