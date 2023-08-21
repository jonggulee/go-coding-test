func solution(t string, p string) int {
	var r int
	for i := 0; i <= len(t)-len(p); i++ {
		if string(t[i:i+len(p)]) <= p {
			r += 1
		}

	}
	return r
}