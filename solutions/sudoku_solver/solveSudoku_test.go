package sudokusolver

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_solveSudoku(t *testing.T) {
	var r [][]byte

	r = [][]byte{
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
	}
	solveSudoku(r)
	assert.Equal(t, [][]byte{
		[]byte("123456789"),
		[]byte("456789123"),
		[]byte("789123456"),
		[]byte("214365897"),
		[]byte("365897214"),
		[]byte("897214365"),
		[]byte("531642978"),
		[]byte("642978531"),
		[]byte("978531642"),
	}, r)

	r = [][]byte{
		[]byte(".23456789"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	solveSudoku(r)
	assert.Equal(t, [][]byte{
		[]byte(".23456789"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}, r)

	r = [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	solveSudoku(r)
	assert.Equal(t, [][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}, r)
}

func Benchmark_solveSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveSudoku([][]byte{
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
		solveSudoku([][]byte{
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
}
