package binarytreelevelordertraversal

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil))
	assert.Equal(t,
		[][]int{
			{1},
		}, levelOrder(&TreeNode{
			Val: 1,
		}))
	assert.Equal(t,
		[][]int{
			{3},
			{9, 20},
			{15, 7},
		}, levelOrder(&TreeNode{
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
	assert.Equal(t,
		[][]int{
			{3},
			{9, 20},
			{11, 15, 7},
		}, levelOrder(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
				Right: &TreeNode{
					Val: 11,
				},
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
}

func Benchmark_levelOrder(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			levelOrder(nil)
			levelOrder(&TreeNode{
				Val: 1,
			})
			levelOrder(&TreeNode{
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
		}
	})
}
