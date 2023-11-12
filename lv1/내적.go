func solution(a []int, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		res += a[i] * b[i]
	}
	return res
}