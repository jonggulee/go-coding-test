func solution(num_list []int, n int) []int {
	ret := []int{}
	for i := 0; i < len(num_list); i = i + n {
		ret = append(ret, num_list[i])
	}
	return ret
}