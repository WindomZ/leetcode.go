package binarytreelevelordertraversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := make([][]int, 0, 4)
	treeOrder(root, &level, 0)
	return level
}

func treeOrder(root *TreeNode, level *[][]int, depth int) {
	if root == nil {
		return
	}
	if len(*level) <= depth {
		*level = append(*level, make([]int, 0, depth*2))
	}
	(*level)[depth] = append((*level)[depth], root.Val) // level[depth] append the value of node
	treeOrder(root.Left, level, depth+1)                // left traversal
	treeOrder(root.Right, level, depth+1)               // right traversal
}
