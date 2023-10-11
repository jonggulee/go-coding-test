func solution(arr []int) []int {
	var ret []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			ret = append(ret, arr[i])
		}
	}
	return ret
}