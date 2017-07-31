package remove_nth_node_from_end_of_list

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil && n >= 1 {
		fix := head
		run := head
		for ; run != nil && n > 0; n-- { // run move n gap
			run = run.Next
		}
		if run == nil { // n >= count of valid head.Next
			return head.Next
		}
		for run = run.Next; run != nil; run = run.Next {
			fix = fix.Next
		}
		fix.Next = fix.Next.Next // remove the nth node from the end of list
		return head
	}
	return head
}
