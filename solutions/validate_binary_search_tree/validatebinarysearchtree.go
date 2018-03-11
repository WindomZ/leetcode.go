package validatebinarysearchtree

import "math"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, low, high int64) bool {
	if root == nil {
		return true
	}
	val := int64(root.Val)
	if val >= high || val <= low {
		return false
	}
	return validBST(root.Left, low, val) && validBST(root.Right, val, high)
}
