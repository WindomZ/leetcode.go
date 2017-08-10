package nextpermutation

func nextPermutation(nums []int) {
	k := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k = i
			break
		}
	}

	if k == -1 {
		// reverse all array
		for k, l := 0, len(nums)-1; k < l; {
			nums[k], nums[l] = nums[l], nums[k]
			k++
			l--
		}
		return
	}

	l := -1
	for i := len(nums) - 1; i > k; i-- {
		if nums[i] > nums[k] {
			l = i
			break
		}
	}

	// swap nums[k] and nums[l]
	nums[k], nums[l] = nums[l], nums[k]

	// reverse array from k+1 to end
	for k, l = k+1, len(nums)-1; k < l; {
		nums[k], nums[l] = nums[l], nums[k]
		k++
		l--
	}
}
