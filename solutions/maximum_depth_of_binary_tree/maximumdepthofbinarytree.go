package maximumdepthofbinarytree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return treeDepth(root, 1)
}

func treeDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth - 1
	}
	// returns max depth between left and right depth
	return max(treeDepth(root.Left, depth+1), treeDepth(root.Right, depth+1))
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
