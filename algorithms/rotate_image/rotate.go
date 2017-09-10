package rotateimage

// rotate clockwise rotate
// flip the matrix vertically, then flip it symmetrically
// 1 2 3    7 8 9    7 4 1
// 4 5 6 => 4 5 6 => 8 5 2
// 7 8 9    1 2 3    9 6 3
func rotate(matrix [][]int) {
	// flip the matrix vertically
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j], matrix[len(matrix)-i-1][j] =
				matrix[len(matrix)-i-1][j], matrix[i][j] // swap
		}
	}
	// flip it symmetrically
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j] // swap
		}
	}
}
