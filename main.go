package main

import (
	"fmt"
)

func solution(arr []int) []int {
	res := []int{}
	if len(arr) <= 1 {
		return []int{-1}
	}

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != min {
			res = append(res, arr[i])
		}
	}

	return res
}

func main() {
	r := solution([]int{4, 3, 2, 1})
	fmt.Println(r)

	r = solution([]int{10})
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
