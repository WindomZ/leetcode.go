package swapnodesinpairs

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	assert.Empty(t, swapPairs(nil))
	assert.Equal(t, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}, swapPairs(&ListNode{
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
	}))
	assert.Equal(t, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}))
}

func Benchmark_swapPairs(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			swapPairs(nil)
			swapPairs(&ListNode{
				Val: 1,
			})
			swapPairs(&ListNode{
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
			})
			swapPairs(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			})
		}
	})
}
