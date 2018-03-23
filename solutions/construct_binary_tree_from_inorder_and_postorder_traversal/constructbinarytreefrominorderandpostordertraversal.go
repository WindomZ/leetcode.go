package constructbinarytreefrominorderandpostordertraversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// current node in postorder
	root := &TreeNode{
		Val: postorder[len(inorder)-1],
	}

	// find the index of current node in inorder
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(inorder)-1])

	return root
}
