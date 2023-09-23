func solution(arr []int) []int {
	ret := []int{}
	for _, v := range arr {
		if v < 50 && v%2 == 1 {
			v = v * 2
			ret = append(ret, v)
		} else if v >= 50 && v%2 == 0 {
			v = v / 2
			ret = append(ret, v)
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}