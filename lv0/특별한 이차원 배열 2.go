func solution(arr [][]int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] == arr[j][i] {
				res *= 1
			} else {
				res *= 0
			}
		}
	}
	return res
}