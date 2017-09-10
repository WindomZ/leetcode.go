package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if l := len(nums); l >= 4 {
		sort.Ints(nums) // sort nums

		r := make([][]int, 0, l)
		for i := 0; i < l-3; i++ {
			if i == 0 || nums[i] != nums[i-1] { // prevents duplicate the first number
				threeSum(&r, nums[i+1:], nums[i], target-nums[i])
			}
		}
		return r
	}
	return [][]int{}
}

func threeSum(r *[][]int, nums []int, base, target int) {
	for i, l := 0, len(nums); i < l-2; i++ { // move i, the first number
		if i == 0 || nums[i] != nums[i-1] { // prevents duplicate the first number
			for n, j, k := nums[i], i+1, l-1; j < k; {
				if sum := n + nums[j] + nums[k]; sum == target {
					*r = append(*r, []int{base, n, nums[j], nums[k]}) // get one of result
					for k--; k > j && nums[k] == nums[k+1]; k-- {     // move k, the third number, no repeated
					}
					for j++; j < k && nums[j] == nums[j-1]; j++ { // move j, the second number, no repeated
					}
				} else if sum > target {
					k-- // move k, the third number
				} else {
					j++ // move j, the second number
				}
			}
		}
	}
}
