package merge_k_sorted_lists

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if l := len(lists); l >= 2 {
		for i, j := 0, l-1; ; {
			if lists[i] = mergeTwoLists(lists[i], lists[j]); j == 1 {
				return lists[0]
			} else if j--; i >= j {
				j = i
				i = 0
			} else if i+1 < j {
				i++
			}
		}
	} else if l == 1 {
		return lists[0]
	}
	return nil
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
