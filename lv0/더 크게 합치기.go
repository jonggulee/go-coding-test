import "strconv"

func solution(a int, b int) int {
	ab, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	ba, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))

	if ab > ba {
		return ab
	} else {
		return ba
	}
}