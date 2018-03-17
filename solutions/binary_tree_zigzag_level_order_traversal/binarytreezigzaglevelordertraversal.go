package binarytreezigzaglevelordertraversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := make([][]int, 0, 4)
	zigzagOrder(root, &level, 0)
	for depth, nodes := range level {
		if depth%2 != 0 { // right to left traversal
			l := len(nodes)
			for i := 0; i < l/2; i++ {
				nodes[i], nodes[l-i-1] = nodes[l-i-1], nodes[i]
			}
		}
	}
	return level
}

func zigzagOrder(root *TreeNode, level *[][]int, depth int) {
	if root == nil {
		return
	}
	if len(*level) <= depth {
		*level = append(*level, make([]int, 0, depth*2))
	}
	(*level)[depth] = append((*level)[depth], root.Val) // level[depth] append the value of node
	zigzagOrder(root.Left, level, depth+1)              // left traversal
	zigzagOrder(root.Right, level, depth+1)             // right traversal
}
