func solution(n int) [][]int {
	spiral := make([][]int, n)

	for i := range spiral {
		spiral[i] = make([]int, n)
	}

	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	x, y, d := 0, 0, 0

	for i := 1; i <= n*n; i++ {
		spiral[x][y] = i
		nx, ny := x+dx[d], y+dy[d]

		if nx < 0 || ny < 0 || nx >= n || ny >= n || spiral[nx][ny] != 0 {
			d = (d + 1) % 4
			nx = x + dx[d]
			ny = y + dy[d]
		}

		x, y = nx, ny
	}

	return spiral
}