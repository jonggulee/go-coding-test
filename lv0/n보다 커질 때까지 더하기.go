func solution(numbers []int, n int) int {
	ret := 0
	for i := 0; i < len(numbers); i++ {
		ret += numbers[i]
		if n < ret {
			return ret
		}
	}

	return ret
}