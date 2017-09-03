package permutations_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums) // really, can choose not to use it!
	result := make([][]int, 0, 10)
	permuteUniqueRecursive(nums, 0, &result)
	return result
}

func permuteUniqueRecursive(nums []int, pos int, result *[][]int) {
	if pos >= len(nums)-1 {
		num := make([]int, len(nums))
		copy(num, nums)
		*result = append(*result, num)
		return
	}
	var used [20]bool // nums in [-9, +9]
	for i := pos; i < len(nums); i++ {
		if used[nums[i]+10] {
			continue // pass duplicate value
		}
		nums[pos], nums[i] = nums[i], nums[pos]     // swap
		permuteUniqueRecursive(nums, pos+1, result) // recursive
		nums[pos], nums[i] = nums[i], nums[pos]     // reset
		used[nums[i]+10] = true                     // keep unique permutations
	}
}
