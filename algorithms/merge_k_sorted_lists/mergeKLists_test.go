package merge_k_sorted_lists

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	assert.Empty(t, mergeKLists(nil))
	assert.Equal(t, &ListNode{Val: 1}, mergeKLists([]*ListNode{{Val: 1}}))
	assert.Equal(t,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		mergeKLists([]*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
			{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}),
	)
	assert.Equal(t,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 7,
									Next: &ListNode{
										Val: 8,
									},
								},
							},
						},
					},
				},
			},
		},
		mergeKLists([]*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
			{
				Val: 2,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
			{
				Val: 3,
				Next: &ListNode{
					Val: 7,
				},
			},
			{
				Val: 4,
				Next: &ListNode{
					Val: 8,
				},
			},
		}),
	)
}

func Benchmark_mergeKLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists(nil)
		mergeKLists([]*ListNode{{Val: 1}})
		mergeKLists([]*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
			{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		})
		mergeKLists([]*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
			{
				Val: 2,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
			{
				Val: 3,
				Next: &ListNode{
					Val: 7,
				},
			},
			{
				Val: 4,
				Next: &ListNode{
					Val: 8,
				},
			},
		})
	}
}
