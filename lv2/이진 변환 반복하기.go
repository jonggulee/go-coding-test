import (
	"fmt"
)

func solution(s string) []int {
	binary := s
	count := 0
	zeroCount := 0
	for {
		if binary == "1" {
			break
		}

		count++
		oneCount := 0
		for _, v := range binary {
			if v == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		binary = decimalToBinary(oneCount)
	}
	return []int{count, zeroCount}
}

func decimalToBinary(decimal int) string {
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal = decimal / 2
	}
	return binary
}