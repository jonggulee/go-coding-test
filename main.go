package main

import (
	"fmt"
)

func solution(board [][]int, k int) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i+j <= k {
				res += board[i][j]
			}
		}
	}
	return res
}

func main() {
	r := solution([][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 2)
	fmt.Println(r)

	// r = solution([][]int{{57, 192, 534, 2}, {9, 345, 192, 999}})
	// fmt.Println(r)

	// r = solution([][]int{{1, 2}, {3, 4}})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
