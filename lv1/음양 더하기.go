func solution(absolutes []int, signs []bool) int {
	var sum int
	for i, v := range absolutes {
		if signs[i] {
			sum += v
		} else {
			sum -= v
		}
	}
	return sum
}