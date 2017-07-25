package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if isEmptyIntSlice(nums1) || isEmptyIntSlice(nums2) {
		if f, ok := findMedian(nums1); ok {
			return f
		} else if f, ok := findMedian(nums2); ok {
			return f
		}
		panic("Arguments should be non-empty!")
	}

	size := len(nums1) + len(nums2)
	odd := size%2 != 0 // bool, if size is odd.
	offset := size / 2 // middle index.
	if !odd {
		offset -= 1 // middle index if size is even.
	}

	var n1, n2 int
	for step := 1; step <= offset; step++ { // step move to middle index
		if nums1[n1] <= nums2[n2] {
			n1++
		} else {
			n2++
		}

		if len(nums1) == n1 {
			return findOffset(nums2, step-n2, odd)
		} else if len(nums2) == n2 {
			return findOffset(nums1, step-n1, odd)
		}
	}

	// calc the median
	if odd {
		return math.Min(float64(nums1[n1]), float64(nums2[n2]))
	} else if len(nums1) >= n1+2 && nums1[n1+1] <= nums2[n2] {
		return 0.5 * float64(nums1[n1]+nums1[n1+1])
	} else if len(nums2) >= n2+2 && nums2[n2+1] <= nums1[n1] {
		return 0.5 * float64(nums2[n2]+nums2[n2+1])
	}
	return 0.5 * float64(nums1[n1]+nums2[n2])
}

func isEmptyIntSlice(arr []int) bool {
	return arr == nil || len(arr) == 0
}

func findMedian(nums []int) (f float64, ok bool) {
	if nums != nil {
		if l := len(nums); l != 0 {
			ok = true
			if l%2 == 0 {
				f = 0.5 * float64(nums[l/2-1]+nums[l/2])
			} else {
				f = float64(nums[l/2])
			}
		}
	}
	return
}

func findOffset(nums []int, offset int, odd bool) float64 {
	idx := (len(nums) - offset) / 2
	if odd {
		return float64(nums[idx])
	}
	return 0.5 * float64(nums[idx-1]+nums[idx])
}
