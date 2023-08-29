import "strconv"

func solution(n int) int {
	answer := 0
	strN := strconv.Itoa(n)
	for _, str := range strN {
		num, _ := strconv.Atoi(string(str))
		answer += num
	}

	return answer
}