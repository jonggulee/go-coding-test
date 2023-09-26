func solution(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}