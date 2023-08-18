func solution(name []string, yearning []int, photo [][]string) []int {
	score := make(map[string]int)
	result := []int{}
	for i := 0; i < len(name); i++ {
		score[name[i]] = yearning[i]
	}

	for i := 0; i < len(photo); i++ {
		sum := 0
		for j := 0; j < len(photo[i]); j++ {
			sum = sum + score[photo[i][j]]
		}
		result = append(result, sum)
	}
	return result
}