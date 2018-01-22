package reverselinkedlistii

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	head = &ListNode{
		Next: head,
	}
	left := head // create a left node
	for i := m; i > 1; i-- {
		left = left.Next // move to m point node before reversing
	}
	right := left.Next
	for i := m; i < n; i++ {
		move := right.Next
		right.Next = move.Next
		move.Next = left.Next
		left.Next = move
	}
	return head.Next
}
