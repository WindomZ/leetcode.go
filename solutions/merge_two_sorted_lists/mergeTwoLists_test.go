package mergetwosortedlists

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_mergeTwoLists_0(t *testing.T) {
	assert.Empty(t, mergeTwoLists(nil, nil))
	assert.Equal(t, &ListNode{
		Val: 1,
	}, mergeTwoLists(nil, &ListNode{
		Val: 1,
	}))
	assert.Equal(t, &ListNode{
		Val: 1,
	}, mergeTwoLists(&ListNode{
		Val: 1,
	}, nil))
}

func Test_mergeTwoLists_1(t *testing.T) {
	r := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	})
	assert.NotEmpty(t, r)
	assert.Equal(t, 1, r.Val)
	assert.Equal(t, 2, r.Next.Val)
	assert.Equal(t, 3, r.Next.Next.Val)
	assert.Equal(t, 4, r.Next.Next.Next.Val)
	assert.Equal(t, 6, r.Next.Next.Next.Next.Val)
	assert.Empty(t, r.Next.Next.Next.Next.Next)
}

func Test_mergeTwoLists_2(t *testing.T) {
	r := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	})
	assert.NotEmpty(t, r)
	assert.Equal(t, 1, r.Val)
	assert.Equal(t, 2, r.Next.Val)
	assert.Equal(t, 3, r.Next.Next.Val)
	assert.Equal(t, 4, r.Next.Next.Next.Val)
	assert.Equal(t, 5, r.Next.Next.Next.Next.Val)
	assert.Equal(t, 6, r.Next.Next.Next.Next.Next.Val)
	assert.Empty(t, r.Next.Next.Next.Next.Next.Next)
}

func Benchmark_mergeTwoLists(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mergeTwoLists(nil, nil)
			mergeTwoLists(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			}, &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
					},
				},
			})
			mergeTwoLists(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			}, &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
					},
				},
			})
		}
	})
}
