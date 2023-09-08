func solution(my_string string, indices []int) string {
	runes := []rune(my_string)

	removeMap := make(map[int]bool)
	for _, index := range indices {
		removeMap[index] = true
	}

	var result []rune
	for i, v := range runes {
		if !removeMap[i] {
			result = append(result, v)
		}
	}

	return string(result)
}