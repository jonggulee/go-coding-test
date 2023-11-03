func solution(n int) [][]int {
	answer := make([][]int, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]int, n)
		answer[i][i] = 1
	}
	return answer
}