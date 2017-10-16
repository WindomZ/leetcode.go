package rotatelist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	size := 1 // number of nodes
	tail := head

	// get the tail node
	for ; tail.Next != nil; size++ {
		tail = tail.Next
	}
	tail.Next = head // circle the nodes

	if k %= size; k != 0 {
		// the tail node is the (size-k)th node (1st node is head)
		for i := 0; i < size-k; i++ {
			tail = tail.Next
		}
	}

	head = tail.Next
	tail.Next = nil
	return head
}
