package rotatelist

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	var node *ListNode
	assert.Equal(t, node, rotateRight(node, 0))

	for i := 5; i >= 1; i-- {
		node = &ListNode{
			Val:  i,
			Next: node,
		}
	}
	node = rotateRight(node, 2)
	val := []int{4, 5, 1, 2, 3}
	for i := 0; i < 5; i++ {
		if assert.NotEmpty(t, node) {
			assert.Equal(t, val[i], node.Val)
			node = node.Next
		}
	}
}

func Benchmark_rotateRight(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	f := func() (node *ListNode) {
		for i := 5; i >= 1; i-- {
			node = &ListNode{
				Val:  i,
				Next: node,
			}
		}
		return
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rotateRight(nil, 0)

			rotateRight(f(), 2)
		}
	})
}
