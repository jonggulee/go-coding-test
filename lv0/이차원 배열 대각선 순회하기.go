func solution(board [][]int, k int) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i+j <= k {
				res += board[i][j]
			}
		}
	}
	return res
}