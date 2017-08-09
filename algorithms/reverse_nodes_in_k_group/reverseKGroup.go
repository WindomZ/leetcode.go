package reversenodesinkgroup

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head != nil && k >= 2 {
		return _reverseKGroup(head, k)
	}
	return head
}

func _reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for ; curr != nil && count != k; count++ { // find the k+1 node
		curr = curr.Next
	}
	if count != k { // if k+1 node is not found
		return head
	}
	if curr != nil {
		curr = _reverseKGroup(curr, k) // reverses list with k+1 node as head
	}
	// curr is k+1 node, count equal k+1
	for ; count > 0; count-- { // reverses current k-group
		next := head.Next // get the next head
		head.Next = curr  // moves curr to the next head
		curr = head       // moves head to curr
		head = next       // moves the next node to head
	}
	return curr
}
