func solution(num_list []int) int {
	var sum int
	var mul = 1
	for i := 0; i < len(num_list); i++ {
		mul *= num_list[i]
		sum += num_list[i]
	}
	sqaureSum := sum * sum
	if mul < sqaureSum {
		return 1
	}

	return 0
}