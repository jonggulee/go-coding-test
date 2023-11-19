func solution(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}

	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}