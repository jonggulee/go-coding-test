func solution(n int) int {
	return n/7 + map[bool]int{true: 1, false: 0}[n%7 > 0]
}