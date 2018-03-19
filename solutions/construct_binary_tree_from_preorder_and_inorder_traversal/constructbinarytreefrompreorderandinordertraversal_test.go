package constructbinarytreefrompreorderandinordertraversal

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_buildTree(t *testing.T) {
	assert.Empty(t, buildTree([]int{}, []int{}))
	assert.Equal(t, &TreeNode{
		Val: 1,
	}, buildTree([]int{1}, []int{1}))
	assert.Equal(t, &TreeNode{
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
	}, buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

func Benchmark_buildTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buildTree([]int{}, []int{})
			buildTree([]int{1}, []int{1})
			buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
		}
	})
}
