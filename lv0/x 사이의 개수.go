func solution(myString string) []int {
	ret := []int{}
	splitX := strings.Split(myString, "x")
	for i := 0; i < len(splitX); i++ {
		count := len(splitX[i])
		ret = append(ret, count)
	}
	return ret
}