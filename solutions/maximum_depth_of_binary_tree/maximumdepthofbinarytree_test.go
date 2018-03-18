package maximumdepthofbinarytree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	assert.Equal(t, 0, maxDepth(nil))
	assert.Equal(t, 1, maxDepth(&TreeNode{Val: 0}))
	assert.Equal(t, 2, maxDepth(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
		},
	}))
	assert.Equal(t, 3, maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
	assert.Equal(t, 5, maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 23,
					Right: &TreeNode{
						Val: 28,
					},
				},
			},
		},
	}))
}

func Benchmark_maxDepth(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			maxDepth(nil)
			maxDepth(&TreeNode{Val: 0})
			maxDepth(&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
			})
			maxDepth(&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			})
			maxDepth(&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 23,
							Right: &TreeNode{
								Val: 28,
							},
						},
					},
				},
			})
		}
	})
}
