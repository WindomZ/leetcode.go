package convertsortedarraytobinarysearchtree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_sortedArrayToBST(t *testing.T) {
	assert.Empty(t, sortedArrayToBST(nil))
	assert.Empty(t, sortedArrayToBST([]int{}))
	assert.Equal(t, &TreeNode{
		Val: 1,
	}, sortedArrayToBST([]int{1}))
	assert.Equal(t, &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}, sortedArrayToBST([]int{1, 2}))
	assert.Equal(t, &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}, sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	assert.Equal(t, &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -10,
			Left: &TreeNode{
				Val: -11,
			},
			Right: &TreeNode{
				Val: -3,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 12,
			},
		},
	}, sortedArrayToBST([]int{-11, -10, -3, 0, 5, 9, 12}))
}

func Benchmark_sortedArrayToBST(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sortedArrayToBST([]int{})
			sortedArrayToBST([]int{1})
			sortedArrayToBST([]int{1, 2})
			sortedArrayToBST([]int{-10, -3, 0, 5, 9})
			sortedArrayToBST([]int{-11, -10, -3, 0, 5, 9, 12})
		}
	})
}
