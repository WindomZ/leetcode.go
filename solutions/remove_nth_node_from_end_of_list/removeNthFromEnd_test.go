package removenthnodefromendoflist

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_removeNthFromEnd_0(t *testing.T) {
	assert.Empty(t, removeNthFromEnd(nil, 0))
	assert.Empty(t, removeNthFromEnd(nil, 1))
	assert.Empty(t, removeNthFromEnd(&ListNode{}, 1))
	assert.Empty(t, removeNthFromEnd(&ListNode{}, 0))
}

func Test_removeNthFromEnd_1(t *testing.T) {
	r := removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, 3)
	assert.NotEmpty(t, r)
	assert.Equal(t, 2, r.Val)
	assert.Equal(t, 3, r.Next.Val)
	assert.Empty(t, r.Next.Next)
}

func Test_removeNthFromEnd_2(t *testing.T) {
	r := removeNthFromEnd(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, 2)
	assert.NotEmpty(t, r)
	assert.Equal(t, 2, r.Val)
	assert.Equal(t, 3, r.Next.Val)
	assert.Empty(t, r.Next.Next)
}

func Test_removeNthFromEnd_3(t *testing.T) {
	r := removeNthFromEnd(&ListNode{
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
	}, 2)
	assert.NotEmpty(t, r)
	assert.Equal(t, 1, r.Val)
	assert.Equal(t, 2, r.Next.Val)
	assert.Equal(t, 3, r.Next.Next.Val)
	assert.Equal(t, 5, r.Next.Next.Next.Val)
	assert.Empty(t, r.Next.Next.Next.Next)
}

func Benchmark_removeNthFromEnd(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			removeNthFromEnd(nil, 0)
			removeNthFromEnd(&ListNode{}, 0)
			removeNthFromEnd(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			}, 3)
			removeNthFromEnd(&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			}, 2)
			removeNthFromEnd(&ListNode{
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
			}, 2)
			removeNthFromEnd(&ListNode{
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
											Next: &ListNode{
												Val: 9,
											},
										},
									},
								},
							},
						},
					},
				},
			}, 1)
		}
	})
}
