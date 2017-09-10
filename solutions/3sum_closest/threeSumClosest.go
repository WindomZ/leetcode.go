package threesumclosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	if l := len(nums); l >= 3 {
		sort.Ints(nums) // sort nums

		r := nums[0] + nums[1] + nums[2]
		for i := 0; i < l-2; i++ { // move i, the first number
			if i != 0 && nums[i] == nums[i-1] { // the first number no repeated
				continue
			}
			for n, j, k := nums[i], i+1, l-1; j < k; {
				sum := n + nums[j] + nums[k]
				diff := abs(target - sum)
				if diff == 0 {
					return target
				} else if sum > target {
					for k--; k > j && nums[k] == nums[k+1]; k-- { // move k, the third number, no repeated
					}
				} else {
					for j++; j < k && nums[j] == nums[j-1]; j++ { // move j, the second number, no repeated
					}
				}
				if diff < abs(target-r) {
					r = sum // get the closest sum
				}
			}
		}

		return r
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
