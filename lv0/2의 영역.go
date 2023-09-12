func solution(arr []int) []int {
	firstIndex := -1
	lastIndex := -1

	for i, v := range arr {
		if v == 2 {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	if lastIndex == -1 {
		return []int{-1}
	}

	return arr[firstIndex : lastIndex+1]
}