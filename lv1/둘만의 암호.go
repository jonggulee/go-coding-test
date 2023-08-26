func solution(s string, skip string, index int) string {
	result := ""
	skipSet := make(map[rune]bool)

	for _, c := range skip {
		skipSet[c] = true
	}

	for _, c := range s {
		nextChar := c
		skipcount := 0

		for i := 0; i < index+skipcount; i++ {
			nextChar++

			if nextChar > 'z' {
				nextChar = 'a'
			}

			if skipSet[nextChar] {
				skipcount++
			}
		}
		result += string(nextChar)
	}

	return result
}