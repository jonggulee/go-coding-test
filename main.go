package main

import (
	"fmt"
)

func solution(n int) int {
	for i := 1; i <= n; i++ {
		if n%i == 1 {
			return i
		}
	}
	return 0
}

func main() {
	r := solution(10)
	fmt.Println(r)

	r = solution(12)
	fmt.Println(r)

	// r = solution(123456, 7)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
