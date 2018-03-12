package sametree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	assert.True(t, isSameTree(nil, nil))
	assert.False(t, isSameTree(&TreeNode{Val: 0}, nil))
	assert.True(t, isSameTree(&TreeNode{Val: 0}, &TreeNode{Val: 0}))
	assert.True(t, isSameTree(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}, &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
	assert.False(t, isSameTree(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}, &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}))
	assert.False(t, isSameTree(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}, &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}))
}

func Benchmark_isSameTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isSameTree(nil, nil)
			isSameTree(&TreeNode{Val: 0}, nil)
			isSameTree(&TreeNode{Val: 0}, &TreeNode{Val: 0})
			isSameTree(&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			}, &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			})
			isSameTree(&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}, &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			})
			isSameTree(&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			}, &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			})
		}
	})
}
