package main

import (
	"fmt"
)

func solution(num int, k int) int {
	for i := 0; i < k; i++ {
		if num[i] == k {
			return i
		}
	}
	return -1
}

func main() {
	r := solution(29183, 1)
	fmt.Println(r)

	// r = solution([][]int{{57, 192, 534, 2}, {9, 345, 192, 999}})
	// fmt.Println(r)

	// r = solution([][]int{{1, 2}, {3, 4}})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
