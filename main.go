package main

import (
	"fmt"
)

func solution(num_list []int, n int) []int {
	ret := []int{}
	for i := 0; i < len(num_list); i = i + n {
		ret = append(ret, num_list[i])
	}
	return ret
}

func main() {
	r := solution([]int{4, 2, 6, 1, 7, 6}, 2)
	fmt.Println(r)

	r = solution([]int{4, 2, 6, 1, 7, 6}, 4)
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
