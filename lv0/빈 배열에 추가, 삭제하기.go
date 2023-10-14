func solution(arr []int, flag []bool) []int {
	ret := []int{}
	for i, v := range flag {
		if v {
			for j := 0; j < arr[i]*2; j++ {
				ret = append(ret, arr[i])
			}
		} else {
			ret = ret[:len(ret)-arr[i]]
		}
	}
	return ret
}