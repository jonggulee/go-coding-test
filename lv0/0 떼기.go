func solution(n_str string) string {
	for {
		if n_str[0] == '0' {
			n_str = n_str[1:]
		} else {
			break
		}
	}

	return n_str
}