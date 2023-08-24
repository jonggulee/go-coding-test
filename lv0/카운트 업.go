func solution(start_num int, end_num int) []int {
	var r []int
	for i := start_num; i <= end_num; i++ {
		r = append(r, i)
	}
	return r
}