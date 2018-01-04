package partitionlist

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// separate the list into left & right distinct lists and link them afterwards.
	left, right := new(ListNode), new(ListNode)
	l, r := left, right
	for ; head != nil; head = head.Next {
		// separate head
		if head.Val < x {
			l.Next = head
			l = head
		} else {
			r.Next = head
			r = head
		}
	}
	// link them
	r.Next = nil
	l.Next = right.Next
	return left.Next
}
