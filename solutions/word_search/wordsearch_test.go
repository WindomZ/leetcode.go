package wordsearch

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_exist(t *testing.T) {
	assert.Equal(t, false, exist([][]byte{}, ""))
	assert.Equal(t, false, exist([][]byte{}, "A"))
	assert.Equal(t, false, exist([][]byte{{}}, "A"))
	assert.Equal(t, true, exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	assert.Equal(t, true, exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
	assert.Equal(t, false, exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}

func Benchmark_exist(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			exist([][]byte{}, "")
			exist([][]byte{}, "A")
			exist([][]byte{{}}, "A")
			exist([][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCCED")
			exist([][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "SEE")
			exist([][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCB")
		}
	})
}
