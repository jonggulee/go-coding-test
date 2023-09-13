package main

import (
	"fmt"
)

func solution(num_list []int, n int) []int {
	return num_list[n-1:]
}

func main() {
	r := solution([]int{2, 1, 6}, 3)
	fmt.Println(r)

	r = solution([]int{5, 2, 1, 7, 5}, 2)
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
