package maximumsubarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// get maximum value between the sum of the contiguous subarray and the current number
		sum = maxInt(sum+nums[i], nums[i])
		// get the largest sum
		max = maxInt(max, sum)
	}
	return max
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
