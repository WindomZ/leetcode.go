package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	l3 := new(ListNode)
	n1, n2, n3 := l1, l2, l3

	for n1 != nil || n2 != nil {
		if n1 != nil {
			n3.Val += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			n3.Val += n2.Val
			n2 = n2.Next
		}

		carry = int(n3.Val / 10)
		if carry == 0 && n1 == nil && n2 == nil {
			break
		}
		n3.Next = &ListNode{Val: carry}
		n3.Val %= 10

		n3 = n3.Next
	}
	return l3
}
