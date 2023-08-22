package main

import (
	"fmt"
)

func solution(numLog []int) string {
	r := ""
	prev := numLog[0]

	for i := 1; i < len(numLog); i++ {
		diff := numLog[i] - prev

		switch diff {
		case 1:
			r += "w"
		case -1:
			r += "s"
		case 10:
			r += "d"
		case -10:
			r += "a"
		}

		prev = numLog[i]
	}

	return r
}

func main() {
	r := solution([]int{0, 1, 0, 10, 0, 1, 0, 10, 0, -1, -2, -1})
	fmt.Println(r)

	// r = solution("500220839878", "7")
	// fmt.Println(r)

	// r = solution("10203", "15")
	// fmt.Println(r)
}
