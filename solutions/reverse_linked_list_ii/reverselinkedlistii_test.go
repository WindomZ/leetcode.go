package reverselinkedlistii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	assert.Empty(t, reverseBetween(nil, 0, 0))
	assert.Equal(t, &ListNode{
		Val: 5,
	}, reverseBetween(&ListNode{
		Val: 5,
	}, 1, 1))
	assert.Equal(t, &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}, reverseBetween(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, 1, 3))
	assert.Equal(t, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, reverseBetween(&ListNode{
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
	}, 2, 4))
}

func Benchmark_reverseBetween(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			reverseBetween(nil, 0, 0)
			reverseBetween(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			}, 1, 1)
			reverseBetween(&ListNode{
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
			}, 2, 4)
		}
	})
}
