func solution(n int, k int) []int {
	r := []int{}
	for i := k; i <= n; i += k {
		r = append(r, i)
	}
	return r
}