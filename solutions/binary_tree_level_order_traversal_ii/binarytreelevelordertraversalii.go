package binarytreelevelordertraversalii

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	list := make([][]int, 0, 2)
	levelOrder(root, &list, 0)
	return list
}

func levelOrder(root *TreeNode, list *[][]int, level int) {
	if root == nil {
		return
	}

	// reverse the list array
	l := len(*list)
	if level >= l {
		if l > 0 {
			*list = append(*list, (*list)[l-1])
			for i := l - 1; i >= 1; i-- {
				(*list)[i] = (*list)[i-1]
			}
			(*list)[0] = make([]int, 0, level*2)
		} else {
			*list = append(*list, make([]int, 0, level*2))
		}
	}

	// recursion the son nodes
	levelOrder(root.Left, list, level+1)
	levelOrder(root.Right, list, level+1)

	// bottom-up level order traversal
	l = len(*list) - level - 1
	(*list)[l] = append((*list)[l], root.Val)
}
