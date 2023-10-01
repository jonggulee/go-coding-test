func solution(myString string, pat string) int {
	count := 0
	patLen := len(pat)

	for i := 0; i <= len(myString)-patLen; i++ {
		if myString[i:i+patLen] == pat {
			count++
		}
	}
	return count
}