package recoverbinarysearchtree

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		},
	}
	recoverTree(root)
	assert.Equal(t, &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
		},
	}, root)

	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	recoverTree(root)
	assert.Equal(t, &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, root)

	root = &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	recoverTree(root)
	assert.Equal(t, &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}, root)

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	recoverTree(root)
	assert.Equal(t, &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}, root)

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			}},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 7,
			}},
	}
	recoverTree(root)
	assert.Equal(t, &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			}},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			}},
	}, root)
}

func Benchmark_recoverTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			recoverTree(&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			})
			recoverTree(&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			})
			recoverTree(&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					}},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					}},
			})
		}
	})
}
