package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target <= 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := make([][]int, 0, 2)
	combinationSum2R(candidates, target, &result, make([]int, 0, len(candidates)), 0, 0)
	return result
}

func combinationSum2R(candidates []int, target int, result *[][]int, stack []int, index, start int) {
	if target == 0 {
		// clone stack[:index]
		temp := make([]int, index)
		for i := 0; i < index; i++ {
			temp[i] = stack[i]
		}
		// append to result
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		if i == start || candidates[i] != candidates[i-1] { // ignore duplicate combination
			stack = append(stack, candidates[i])                                            // push
			combinationSum2R(candidates, target-candidates[i], result, stack, index+1, i+1) // next
			stack = stack[:index]                                                           // pop
		}
	}
}
