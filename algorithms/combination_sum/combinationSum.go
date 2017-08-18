package combinationsum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target <= 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := make([][]int, 0, 2)
	combinationSumR(candidates, target, &result, make([]int, 0, len(candidates)), 0, 0)
	return result
}

func combinationSumR(candidates []int, target int, result *[][]int, stack []int, index, start int) {
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
		stack = append(stack, candidates[i])                                         // push
		combinationSumR(candidates, target-candidates[i], result, stack, index+1, i) // next
		stack = stack[:index]                                                        // pop
	}
}
