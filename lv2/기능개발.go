func solution(progresses []int, speeds []int) []int {
	var result []int
	var completeDays []int

	for i := 0; i < len(progresses); i++ {
		days := (100 - progresses[i] + speeds[i] - 1) / speeds[i]
		completeDays = append(completeDays, days)
	}

	for len(completeDays) > 0 {
		first := completeDays[0]
		count := 0

		for len(completeDays) > 0 && completeDays[0] <= first {
			count++
			completeDays = completeDays[1:]
		}

		result = append(result, count)
	}

	return result
}