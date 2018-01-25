package binarytreeinordertraversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		// push all the left children from root node.
		result = append(result, inorderTraversal(root.Left)...)
		// push the root node
		result = append(result, root.Val)
		// push all the right children from root node.
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
