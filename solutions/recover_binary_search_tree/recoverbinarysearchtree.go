package recoverbinarysearchtree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var pre, cur, first, second *TreeNode
	// Morris Traversal
	// https://en.wikipedia.org/wiki/Threaded_binary_tree#The_array_of_Inorder_traversal
	for root != nil {
		if root.Left != nil {
			// connect threading for root
			cur = root.Left
			for cur.Right != nil && cur.Right != root {
				cur = cur.Right
			}
			// the threading already exists
			if cur.Right != nil {
				if pre != nil && pre.Val > root.Val {
					// found the mistaken nodes
					if first == nil {
						first = pre
					}
					second = root
				}
				pre, cur.Right, root = root, nil, root.Right
			} else {
				// construct the threading
				cur.Right, root = root, root.Left
			}
		} else {
			if pre != nil && pre.Val > root.Val {
				// found the mistaken nodes
				if first == nil {
					first = pre
				}
				second = root
			}
			pre, root = root, root.Right
		}
	}
	// swap the two mistaken nodes
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
