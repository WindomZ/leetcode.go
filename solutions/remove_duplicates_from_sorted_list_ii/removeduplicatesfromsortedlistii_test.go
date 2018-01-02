package removeduplicatesfromsortedlistii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	assert.Empty(t, deleteDuplicates(nil))
	assert.Equal(t, &ListNode{
		Val: 0,
	}, deleteDuplicates(&ListNode{
		Val: 0,
	}))
	assert.Equal(t, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}))
	assert.Equal(t,
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		}, deleteDuplicates(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}))
	assert.Equal(t,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
				},
			},
		}, deleteDuplicates(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		}))
}

func Benchmark_deleteDuplicates(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			deleteDuplicates(nil)
			deleteDuplicates(&ListNode{
				Val: 0,
			})
			deleteDuplicates(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			})
			deleteDuplicates(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			})
			deleteDuplicates(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			})
			deleteDuplicates(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
			})
		}
	})
}
