package permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0, 2*len(nums))
	permuteRecursive(nums, 0, &result)
	return result
}

func permuteRecursive(nums []int, pos int, result *[][]int) {
	if pos >= len(nums) {
		num := make([]int, len(nums))
		copy(num, nums)
		*result = append(*result, num) // append to result
		return
	}
	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos] // swap
		permuteRecursive(nums, pos+1, result)   // recursive
		nums[pos], nums[i] = nums[i], nums[pos] // reset
	}
}
