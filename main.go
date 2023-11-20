package main

import (
	"fmt"
)

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	res := make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = make([]int, len(arr1[i]))
		for j := 0; j < len(arr1[i]); j++ {
			res[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return res
}

func main() {
	r := solution([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}})
	fmt.Println(r)

	r = solution([][]int{{1}, {2}}, [][]int{{3}, {4}})
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
