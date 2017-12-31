package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	if len(nums) >= 3 {
		r := 2 // at least 2 difference
		for i := 2; i < len(nums); i++ {
			if nums[i] != nums[r-2] {
				nums[r] = nums[i] // remove the duplicates
				r++               // count from different values
			}
		}
		return r
	}
	return len(nums)
}
