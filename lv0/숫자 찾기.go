import "strconv"

func solution(num int, k int) int {
	strNum := strconv.Itoa(num)
	strK := strconv.Itoa(k)

	for i, v := range strNum {
		if string(v) == strK {
			return i + 1
		}
	}

	return -1
}