package main

import (
	"fmt"
)

func solution(a []int, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		res += a[i] * b[i]
	}
	return res
}

func main() {
	r := solution([]int{1, 2, 3, 4}, []int{-3, -1, 0, 2})
	fmt.Println(r)

	r = solution([]int{-1, 0, 1}, []int{1, 0, -1})
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
