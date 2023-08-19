import "strconv"

func solution(a int, b int) int {
	concatenatedSum, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	productSum := 2 * a * b

	if concatenatedSum > productSum {
		return concatenatedSum
	} else {
		return productSum
	}
}