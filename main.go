package main

import (
	"fmt"
)

func solution(n int) [][]int {
	answer := make([][]int, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]int, n)
		answer[i][i] = 1
	}
	return answer
}

func main() {
	r := solution(3)
	fmt.Println(r)

	r = solution(6)
	fmt.Println(r)

	r = solution(1)
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
