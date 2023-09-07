func solution(my_string string) []int {
	r := make([]int, 52)
	for _, ch := range my_string {
		if 'A' <= ch && ch <= 'Z' {
			index := ch - 'A'
			r[index]++
		} else {
			index := ch - 'a' + 26
			r[index]++
		}
	}

	return r
}