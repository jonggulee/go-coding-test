package main

import (
	"fmt"
)

func solution(arr []int, queries [][]int) []int {
	for _, query := range queries {
		s, e, k := query[0], query[1], query[2]
		for i := s; i <= e; i++ {
			if i%k == 0 {
				arr[i] += 1
			}
		}
	}
	return arr
}

func main() {
	r := solution([]int{0, 1, 2, 4, 3}, [][]int{{0, 4, 1}, {0, 3, 2}, {0, 3, 3}})
	fmt.Println(r)

	// r = solution("78720646226947352489")
	// fmt.Println(r)
}
