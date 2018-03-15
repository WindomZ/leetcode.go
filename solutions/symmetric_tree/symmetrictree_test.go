package symmetrictree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	assert.True(t, isSymmetric(nil))
	assert.True(t, isSymmetric(&TreeNode{Val: 0}))
	assert.True(t, isSymmetric(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
	}))
	assert.False(t, isSymmetric(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 1},
		},
	}))
	assert.False(t, isSymmetric(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 1},
		},
	}))
}

func Benchmark_isSymmetric(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isSymmetric(nil)
			isSymmetric(&TreeNode{Val: 0})
			isSymmetric(&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			})
			isSymmetric(&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 1},
				},
			})
			isSymmetric(&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 1},
				},
			})
		}
	})
}
