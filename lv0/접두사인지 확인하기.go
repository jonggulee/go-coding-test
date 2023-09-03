func solution(my_string string, is_prefix string) int {
	for i := 0; i < len(my_string); i++ {
		if my_string[:i] == is_prefix {
			return 1
		}
	}
	return 0
}