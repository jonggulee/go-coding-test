func solution(str_list []string) []string {
	for i, v := range str_list {
		if v == "l" {
			return str_list[:i]
		} else if v == "r" {
			return str_list[i+1:]
		}
	}
	return []string{}
}