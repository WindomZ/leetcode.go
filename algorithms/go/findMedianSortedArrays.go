package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if isEmptyIntSlice(nums1) || isEmptyIntSlice(nums2) {
		if f, ok := findMedian(nums1); ok {
			return f
		} else if f, ok := findMedian(nums2); ok {
			return f
		}
		panic("Should be non-empty parameters")
	}

	size := len(nums1) + len(nums2)
	if size%2 != 0 {
		return sortedMin(nums1, nums2, size/2, true)
	}
	return sortedMin(nums1, nums2, size/2-1, false)
}

func isEmptyIntSlice(arr []int) bool {
	return arr == nil || len(arr) == 0
}

func sortedMin(nums1 []int, nums2 []int, offset int, odd bool) float64 {
	//println(fmt.Sprintf("%v %v %v %v", nums1, nums2, offset, odd))
	if offset == 0 {
		if isEmptyIntSlice(nums1) {
			if odd {
				return float64(nums2[0])
			}
			return 0.5 * float64(nums2[0]+nums2[1])
		} else if isEmptyIntSlice(nums2) {
			if odd {
				return float64(nums1[0])
			}
			return 0.5 * float64(nums1[0]+nums1[1])
		}

		if odd {
			return math.Min(float64(nums1[0]), float64(nums2[0]))
		} else if len(nums1) >= 2 && nums1[1] <= nums2[0] {
			return 0.5 * float64(nums1[0]+nums1[1])
		} else if len(nums2) >= 2 && nums2[1] <= nums1[0] {
			return 0.5 * float64(nums2[0]+nums2[1])
		}
		return 0.5 * float64(nums1[0]+nums2[0])
	}

	if isEmptyIntSlice(nums1) {
		return sortedMin(nil, nums2[1:], offset-1, odd)
	} else if isEmptyIntSlice(nums2) {
		return sortedMin(nums1[1:], nil, offset-1, odd)
	}

	if nums1[0] <= nums2[0] {
		if len(nums1) >= 2 {
			return sortedMin(nums1[1:], nums2, offset-1, odd)
		}
		return sortedMin(nil, nums2, offset-1, odd)
	}
	if len(nums2) >= 2 {
		return sortedMin(nums1, nums2[1:], offset-1, odd)
	}
	return sortedMin(nums1, nil, offset-1, odd)
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
