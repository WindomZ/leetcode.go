package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	spiral := make([]int, len(matrix)*len(matrix[0]))
	for up, down, left, right, i := 0, len(matrix)-1, 0, len(matrix[0])-1, 0; ; {
		// up
		for col := left; col <= right; col++ {
			spiral[i] = matrix[up][col]
			i++
		}
		if up++; up > down {
			break
		}
		// right
		for row := up; row <= down; row++ {
			spiral[i] = matrix[row][right]
			i++
		}
		if right--; right < left {
			break
		}
		// down
		for col := right; col >= left; col-- {
			spiral[i] = matrix[down][col]
			i++
		}
		if down--; down < up {
			break
		}
		// left
		for row := down; row >= up; row-- {
			spiral[i] = matrix[row][left]
			i++
		}
		if left++; left > right {
			break
		}
	}
	return spiral
}
