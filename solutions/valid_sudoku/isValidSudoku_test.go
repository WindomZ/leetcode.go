package validsudoku

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_isValidSudoku(t *testing.T) {
	assert.Equal(t, true, isValidSudoku([][]byte{
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
	}))
	assert.Equal(t, true, isValidSudoku([][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}))
	assert.Equal(t, true, isValidSudoku([][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}))
	assert.Equal(t, false, isValidSudoku([][]byte{
		[]byte("53..75..."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}))
}

func Benchmark_isValidSudoku(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isValidSudoku([][]byte{
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
				[]byte("........."),
			})
			isValidSudoku([][]byte{
				[]byte(".87654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			})
			isValidSudoku([][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			})
		}
	})
}
