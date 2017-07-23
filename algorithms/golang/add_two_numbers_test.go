package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_addTwoNumbers1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	r := addTwoNumbers(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}

func Test_addTwoNumbers2(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
	}
	l2 := &ListNode{
		Val: 5,
	}
	l3 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}

	r := addTwoNumbers(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}

func Test_addTwoNumbers3(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
		},
	}
	l2 := &ListNode{
		Val: 0,
	}
	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
		},
	}

	r := addTwoNumbers(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}

func Test_addTwoNumbers4(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 8,
		},
	}
	l2 := &ListNode{
		Val: 1,
	}
	l3 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 9,
		},
	}

	r := addTwoNumbers(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
