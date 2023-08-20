func solution(ineq string, eq string, n int, m int) int {
	var cond bool

	switch ineq + eq {
	case "<=":
		cond = n <= m
	case "<!":
		cond = n < m
	case ">=":
		cond = n >= m
	case ">!":
		cond = n > m
	}

	if cond {
		return 1
	}
	return 0
}