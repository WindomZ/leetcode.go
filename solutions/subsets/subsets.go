package subsets

// like: {1, 2, 3}
// 0) 0 0 0  -> {}
// 1) 0 0 1  -> {1}
// 2) 0 1 0  -> {2}
// 3) 0 1 1  -> {1,2}
// 4) 1 0 0  -> {3}
// 5) 1 0 1  -> {1,3}
// 6) 1 1 0  -> {2,3}
// 7) 1 1 1  -> {1,2,3}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	count := 2
	for i := 1; i < len(nums); i++ {
		count *= 2 // count of result is 2^N
	}
	result := make([][]int, count)
	result[0] = []int{}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < count; j++ {
			if (uint(j)>>uint(i))&1 != 0 { // 2^N possible outcomes
				result[j] = append(result[j], nums[i])
			}
		}
	}
	return result
}
