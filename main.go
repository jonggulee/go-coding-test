package main

import (
	"fmt"
)

func solution(num_list []int, n int) int {
	for _, v := range num_list {
		if v == n {
			return 1
		}
	}
	return 0
}

func main() {
	r := solution([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(r)

	r = solution([]int{15, 98, 23, 2, 15}, 20)
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
