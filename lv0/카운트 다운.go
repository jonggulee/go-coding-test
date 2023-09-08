func solution(start int, end_num int) []int {
	var ret []int
	for i := start; i >= end_num; i-- {
		ret = append(ret, i)
	}
	return ret
}