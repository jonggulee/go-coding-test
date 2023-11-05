package main

import (
	"fmt"
)

func solution(arr [][]int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] == arr[j][i] {
				res *= 1
			} else {
				res *= 0
			}
		}
	}
	return res
}

func main() {
	r := solution([][]int{{5, 192, 33}, {192, 72, 95}, {33, 95, 999}})
	fmt.Println(r)

	r = solution([][]int{{19, 498, 258, 587}, {3, 93, 7, 754}, {258, 7, 1000, 723}, {587, 754, 723, 81}})
	fmt.Println(r)

	// r = solution(1)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
