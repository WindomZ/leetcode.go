package binarytreelevelordertraversalii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_levelOrderBottom(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrderBottom(nil))
	assert.Equal(t, [][]int{{1}}, levelOrderBottom(&TreeNode{
		Val: 1,
	}))
	assert.Equal(t, [][]int{
		{15, 7},
		{9, 20},
		{3},
	}, levelOrderBottom(&TreeNode{
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
}

func Benchmark_levelOrderBottom(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			levelOrderBottom(nil)
			levelOrderBottom(&TreeNode{
				Val: 1,
			})
			levelOrderBottom(&TreeNode{
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
