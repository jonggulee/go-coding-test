func solution(myString string, pat string) int {
	lowString := strings.ToLower(myString)
	lowPat := strings.ToLower(pat)

	if strings.Contains(lowString, lowPat) {
		return 1
	}

	return 0
}