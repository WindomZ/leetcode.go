package swapnodesinpairs

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		next := head.Next
		head.Next = swapPairs(head.Next.Next) // swaps others from next pairs
		next.Next = head                      // swaps this pairs: head and next=head.Next
		return next
	}
	return head
}
