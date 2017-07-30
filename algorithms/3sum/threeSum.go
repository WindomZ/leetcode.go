package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	if l := len(nums); l >= 3 {
		sort.Ints(nums) // sort nums

		r := make([][]int, 0, l)
		for i := 0; i < l-2; i++ { // move i, the first number
			if i != 0 && nums[i] == nums[i-1] { // the first number no repeated
				continue
			}
			for n, j, k := nums[i], i+1, l-1; j < k; {
				if sum := n + nums[j] + nums[k]; sum == 0 {
					r = append(r, []int{n, nums[j], nums[k]})     // get one of result
					for k--; k > j && nums[k] == nums[k+1]; k-- { // move k, the third number, no repeated
					}
					for j++; j < k && nums[j] == nums[j-1]; j++ { // move j, the second number, no repeated
					}
				} else if sum > 0 {
					k-- // move k, the third number
				} else {
					j++ // move j, the second number
				}
			}
		}

		return r
	}
	return [][]int{}
}
