func solution(num_list []int) int {
	if len(num_list) >= 11 {
		sum := 0
		for _, num := range num_list {
			sum += num
		}
		return sum
	} else {
		product := 1
		for _, num := range num_list {
			product *= num
		}
		return product
	}
}