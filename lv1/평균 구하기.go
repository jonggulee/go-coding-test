func solution(arr []int) float64 {
	sum := 0
	lenArr := len(arr)
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(lenArr)
}