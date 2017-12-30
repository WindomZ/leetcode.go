package wordsearch

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if isExist(board, word, m, n, x, y, 0) {
				return true
			}
		}
	}
	return false
}

func isExist(board [][]byte, word string, m, n, x, y, z int) bool {
	if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != word[z] {
		return false
	}
	t := board[x][y]
	board[x][y] = 0
	if z++; z >= len(word) ||
		isExist(board, word, m, n, x-1, y, z) ||
		isExist(board, word, m, n, x+1, y, z) ||
		isExist(board, word, m, n, x, y-1, z) ||
		isExist(board, word, m, n, x, y+1, z) {
		return true
	}
	board[x][y] = t
	return false
}
