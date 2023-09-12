func solution(arr []int, intervals [][]int) []int {
	ret := []int{}
	for _, v := range intervals {
		ret = append(ret, arr[v[0]:v[1]+1]...)
	}
	return ret
}