package sortcolors

func sortColors(nums []int) {
	i1, i2, i3 := 0, 0, 0
	for _, v := range nums {
		if v == 0 { // insert 0
			nums[i3] = 2
			nums[i2] = 1
			nums[i1] = 0
			i1++
			i2++
			i3++
		} else if v == 1 { // insert 1
			nums[i3] = 2
			nums[i2] = 1
			i2++
			i3++
		} else if v == 2 { // insert 2
			nums[i3] = 2
			i3++
		}
	}
}
