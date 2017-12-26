package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	l, r := 0, m*n-1
	for l != r {
		mid := (l + r - 1) >> 1 // mid is the middle index
		if matrix[mid/m][mid%m] < target {
			l = mid + 1 // move to right
		} else {
			r = mid // move to left
		}
	}
	return matrix[r/m][r%m] == target
}
