func solution(n int) []int {
	var r []int
	r = append(r, n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			r = append(r, n)
		} else {
			n = 3*n + 1
			r = append(r, n)
		}
	}
	return r
}