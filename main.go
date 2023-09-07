package main

import (
	"fmt"
)

func solution(n int, k int) []int {
	r := []int{}
	for i := k; i <= n; i += k {
		r = append(r, i)
	}
	return r
}

func main() {
	r := solution(10, 3)
	fmt.Println(r)

	r = solution(15, 5)
	fmt.Println(r)

	// r = solution(-4, 2)
	// fmt.Println(r)
}
