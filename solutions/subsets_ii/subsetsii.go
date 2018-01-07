package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums) // sort nums

	count := 2
	for i := 1; i < len(nums); i++ {
		count *= 2 // count of max result is 2^N
	}
	result := make([][]int, 1, count)
	result[0] = []int{} // first is {}

	for i, l, r := 0, 0, 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			l = r // offset l if the next number is repeated
		} else {
			l = 0 // copy and append the next number from the previous result
		}
		r = len(result) // move r to end of result
		for j := l; j < r; j++ {
			// copy new point
			t := make([]int, len(result[j])+1)
			copy(t, result[j])
			t[len(t)-1] = nums[i]
			// double the result and append the next number
			result = append(result, t)
		}
	}
	return result
}
