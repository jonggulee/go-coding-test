package main

import (
	"fmt"
	"strconv"
)

func solution(n int) string {
	return strconv.Itoa(n)
}

func main() {
	r := solution(123)
	fmt.Println(r)

	r = solution(2573)
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
