func solution(my_string string, is_suffix string) int {
	for i := 0; i < len(my_string); i++ {
		if my_string[i:] == is_suffix {
			return 1
		}
	}
	return 0
}