package main

import (
	"fmt"
)

func solution(arr []int, delete_list []int) []int {
	res := []int{}
	deleteMap := make(map[int]bool)

	for _, d := range delete_list {
		deleteMap[d] = true
	}

	for _, v := range arr {
		if !deleteMap[v] {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	r := solution([]int{293, 1000, 395, 678, 94}, []int{94, 777, 104, 1000, 1, 12})
	fmt.Println(r)

	r = solution([]int{110, 66, 439, 785, 1}, []int{377, 823, 119, 43})
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
