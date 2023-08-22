func solution(numLog []int) string {
	r := ""
	prev := numLog[0]

	for i := 1; i < len(numLog); i++ {
		diff := numLog[i] - prev

		switch diff {
		case 1:
			r += "w"
		case -1:
			r += "s"
		case 10:
			r += "d"
		case -10:
			r += "a"
		}

		prev = numLog[i]
	}

	return r
}