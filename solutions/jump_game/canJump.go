package jumpgame

func canJump(nums []int) bool {
	var i int
	// j is the maximum jump length
	for j := 0; i < len(nums) && i <= j; i++ {
		j = max(i+nums[i], j)
	}
	return i == len(nums)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
