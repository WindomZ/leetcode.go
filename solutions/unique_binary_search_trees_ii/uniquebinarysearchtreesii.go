package uniquebinarysearchtreesii

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return generateTreesSE(1, n)
}

func generateTreesSE(start, end int) []*TreeNode {
	if start > end {
		return nil
	}
	if end == 1 || start == end {
		return []*TreeNode{{ // returns list with a node if only one value
			Val: start,
		}}
	}
	if end == 2 {
		return []*TreeNode{
			{
				Val:   start,
				Right: &TreeNode{Val: start + 1},
			},
			{
				Val:  start + 1,
				Left: &TreeNode{Val: start},
			},
		}
	}

	list := make([]*TreeNode, 0, end-start)
	for i := start; i <= end; i++ {
		left := generateTreesSE(start, i-1) // get the left nodes
		right := generateTreesSE(i+1, end)  // get the right nodes
		if left == nil {
			for _, r := range right {
				list = append(list, &TreeNode{
					Val:   i,
					Right: r,
				})
			}
			continue
		}
		if right == nil {
			for _, l := range left {
				list = append(list, &TreeNode{
					Val:  i,
					Left: l,
				})
			}
			continue
		}
		for _, l := range left {
			for _, r := range right {
				list = append(list, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return list
}
