package main

import (
	"fmt"
)

func solution(num_list []int) int {
	for i, v := range num_list {
		if v < 0 {
			return i
		}
	}
	return -1
}

func main() {
	r := solution([]int{12, 4, 15, 46, 38, -2, 15})
	fmt.Println(r)

	// r = solution(4, []int{1, 5, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// fmt.Println(r)

	// r = solution(24, 24)
	// fmt.Println(r)
}
