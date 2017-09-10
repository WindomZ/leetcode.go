package sudokusolver

func solveSudoku(board [][]byte) {
	b := newBoard()
	for x, cs := range board {
		for y, n := range cs {
			if n != '.' {
				b.checkCell(x, y, n-'0')
			}
		}
	}
	if b = b.pass(); b != nil {
		for x, cs := range b {
			for y, c := range cs {
				if board[x][y] == '.' {
					board[x][y] = c.value() + '0'
				}
			}
		}
	}
}

const cellChecked = cell(1 << 15) // use 15 bit

type cell uint16 // use only 1~9 bits

// checked pass the board check
func (c cell) checked() bool {
	return (c & cellChecked) != 0
}

// value get the unique value
func (c cell) value() byte {
	c &^= cellChecked
	r := cell(1)

	for i := byte(1); i <= 9; i++ {
		if r == c {
			return i
		}
		r <<= 1
	}
	return 0
}

type board [][]cell

func newBoard() board {
	ret := make(board, 9)
	for idx := range ret {
		ret[idx] = make([]cell, 9)
		for i := range ret[idx] {
			ret[idx][i] = (1 << 9) - 1
		}
	}
	return ret
}

func (b board) checkCell(x, y int, v byte) {
	mask := cell(1 << (v - 1))

	// excludes others
	for i := 0; i < 9; i++ {
		b[x][i] &^= mask
		b[i][y] &^= mask
	}
	sx := x / 3 * 3
	sy := y / 3 * 3
	for i := sx; i < sx+3; i++ {
		for j := sy; j < sy+3; j++ {
			b[i][j] &^= mask
		}
	}

	b[x][y] = mask | cellChecked // set cell(x, y) checked
}

func (b board) pass() board {
	for changed := true; changed; {
		changed = false
		for x, cs := range b {
			for y, c := range cs {
				if c == 0 {
					return nil
				} else if c.checked() {
				} else if n := c.value(); n > 0 {
					b.checkCell(x, y, n)
					changed = true
				}
			}
		}
	}
	for x, cs := range b {
		for y, c := range cs {
			if !c.checked() {
				for k := byte(0); k < 9; k++ {
					if (c & (1 << k)) != 0 {
						cb := b.clone()
						cb.checkCell(x, y, k+1)
						if cb = cb.pass(); cb != nil {
							return cb
						}
					}
				}
				return nil
			}
		}
	}
	return b
}

func (b board) clone() board {
	ret := make(board, 9)
	for idx := range ret {
		ret[idx] = make([]cell, 9)
		for i := range ret[idx] {
			ret[idx][i] = b[idx][i]
		}
	}
	return ret
}
