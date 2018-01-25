package binarytreeinordertraversal

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	assert.Empty(t, inorderTraversal(nil))
	assert.Equal(t,
		[]int{1},
		inorderTraversal(&TreeNode{
			Val: 1,
		}))
	assert.Equal(t,
		[]int{1, 3, 2},
		inorderTraversal(&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}))
	assert.Equal(t,
		[]int{5, 3, 1, 2, 4, 6},
		inorderTraversal(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		}))
}

func Benchmark_inorderTraversal(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			inorderTraversal(nil)
			inorderTraversal(&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			})
			inorderTraversal(&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			})
		}
	})
}
