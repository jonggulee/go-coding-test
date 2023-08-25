import "strconv"

func solution(number string) int {
	var sum int
	for i := 0; i < len(number); i++ {
		numberInt, _ := strconv.Atoi(string((number[i])))
		sum += numberInt
	}
	remainder := sum % 9
	return remainder
}