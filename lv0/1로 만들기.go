func solution(num_list []int) int {
	count := 0
	for i := 0; i < len(num_list); i++ {
		for num_list[i] != 1 {
			if num_list[i]%2 == 1 {
				num_list[i] = (num_list[i] - 1) / 2
				count++
			} else if num_list[i]%2 == 0 {
				num_list[i] /= 2
				count++
			}
		}
	}

	return count
}