package graycode

func grayCode(n int) []int {
	result := make([]int, int(1<<uint(n)))
	for i := 0; i < len(result); i++ {
		result[i] = i ^ i>>1
	}
	return result
}
