package main

import (
	"fmt"
)

func solution(arr []int, intervals [][]int) []int {
	ret := []int{}
	for _, v := range intervals {
		ret = append(ret, arr[v[0]:v[1]+1]...)
	}
	return ret
}

func main() {
	r := solution([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 4}})
	fmt.Println(r)

	// r = solution(4, []int{1, 5, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// fmt.Println(r)

	// r = solution(24, 24)
	// fmt.Println(r)
}
