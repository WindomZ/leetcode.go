package removeelement

func removeElement(nums []int, val int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
