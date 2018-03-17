package binarytreezigzaglevelordertraversal

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	assert.Equal(t, [][]int{}, zigzagLevelOrder(nil))
	assert.Equal(t,
		[][]int{
			{1},
		}, zigzagLevelOrder(&TreeNode{
			Val: 1,
		}))
	assert.Equal(t,
		[][]int{
			{3},
			{20, 9},
			{15, 7},
		}, zigzagLevelOrder(&TreeNode{
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
			{20, 9},
			{11, 15, 7},
		}, zigzagLevelOrder(&TreeNode{
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

func Benchmark_zigzagLevelOrder(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zigzagLevelOrder(nil)
			zigzagLevelOrder(&TreeNode{
				Val: 1,
			})
			zigzagLevelOrder(&TreeNode{
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
