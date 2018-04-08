package convertsortedarraytobinarysearchtree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArray(&nums, 0, len(nums)-1)
}

func sortedArray(nums *[]int, low, high int) *TreeNode {
	if low > high { // Done
		return nil
	}
	mid := low + (high-low)/2
	if (high-low)%2 != 0 {
		mid++
	}
	return &TreeNode{
		Val:   (*nums)[mid],
		Left:  sortedArray(nums, low, mid-1),
		Right: sortedArray(nums, mid+1, high),
	}
}
