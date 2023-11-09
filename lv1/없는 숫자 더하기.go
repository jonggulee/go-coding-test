func solution(numbers []int) int {
	res := 45
	for _, v := range numbers {
		res -= v
	}

	return res
}