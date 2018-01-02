package removeduplicatesfromsortedlistii

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	pre := &ListNode{
		Next: head,
	}
	for cur := pre; cur.Next != nil && cur.Next.Next != nil; {
		if cur.Next.Val == cur.Next.Next.Val {
			dummy := cur
			for dummy.Next != nil && dummy.Next.Next != nil &&
				dummy.Next.Val == dummy.Next.Next.Val {
				dummy = dummy.Next
			}
			cur.Next = dummy.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return pre.Next
}
