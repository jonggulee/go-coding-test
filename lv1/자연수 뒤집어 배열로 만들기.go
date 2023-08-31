func solution(n int64) []int {
	var answer []int

	temp := n
	for temp > 0 {
		answer = append(answer, int(temp%10))
		temp /= 10
	}
	return answer
}