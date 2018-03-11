package validatebinarysearchtree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	assert.True(t, isValidBST(nil))
	assert.False(t, isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}))
	assert.True(t, isValidBST(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}))
	assert.True(t, isValidBST(&TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}))
	assert.False(t, isValidBST(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 2},
		},
	}))
	assert.False(t, isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 2},
		},
	}))
	assert.False(t, isValidBST(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
	assert.False(t, isValidBST(&TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}))
	assert.False(t, isValidBST(&TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 1},
		},
	}))
}

func Benchmark_isValidBST(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isValidBST(nil)
			isValidBST(&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
			})
			isValidBST(&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			})
			isValidBST(&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			})
			isValidBST(&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 2},
				},
			})
			isValidBST(&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 2},
				},
			})
			isValidBST(&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			})
			isValidBST(&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			})
			isValidBST(&TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
			})
		}
	})
}
