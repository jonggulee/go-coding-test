func solution(strArr []string) int {
	freqMap := make(map[int]int)
	maxCount := 0
	for _, str := range strArr {
		lenght := len(str)
		freqMap[lenght]++

		if freqMap[lenght] > maxCount {
			maxCount = freqMap[lenght]
		}
	}
	return maxCount
}