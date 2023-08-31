func solution(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ret += n / i
		}
	}
	return ret
}