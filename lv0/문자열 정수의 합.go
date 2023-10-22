func solution(num_str string) int {
	res := 0
	for _, v := range num_str {
		res += int(v - '0')
	}

	return res
}