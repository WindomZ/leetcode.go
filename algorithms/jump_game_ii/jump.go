package jumpgameii

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	for jump, start, end, max := 0, 0, 0, 0; end >= start; end = max {
		for ; start <= end; start++ { // find the max step index from the steps(num[start:end])
			if nums[start]+start > max {
				max = nums[start] + start // find the max step index
			}
			if max >= len(nums)-1 {
				return jump + 1 // finish
			}
		}
		jump++ // next step
	}
	return 0
}
