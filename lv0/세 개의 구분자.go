func solution(myStr string) []string {
	var result []string
	var temp string

	for _, ch := range myStr {
		if ch == 'a' || ch == 'b' || ch == 'c' {
			if len(temp) > 0 {
				result = append(result, temp)
			}
			temp = ""
		} else {
			temp += string(ch)
		}
	}

	if len(temp) > 0 {
		result = append(result, temp)
	}

	if len(result) == 0 {
		return []string{"EMPTY"}
	}

	return result
}