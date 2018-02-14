package uniquebinarysearchtreesii

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_generateTrees(t *testing.T) {
	assert.Equal(t, []*TreeNode{}, generateTrees(0))
	assert.Equal(t,
		[]*TreeNode{
			{
				Val: 1,
			},
		},
		generateTrees(1))
	assert.Equal(t,
		[]*TreeNode{
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		generateTrees(2))
	assert.Equal(t,
		[]*TreeNode{
			{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			{
				Val: 1,
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 2},
				},
			},
			{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
			{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
		},
		generateTrees(3))
}

func Benchmark_generateTrees(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generateTrees(0)
			generateTrees(2)
			generateTrees(3)
		}
	})
}
