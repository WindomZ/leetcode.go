package constructbinarytreefrompreorderandinordertraversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// current node in preorder
	root := &TreeNode{
		Val: preorder[0],
	}

	// find the index of current node in inorder
	index := 0
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
