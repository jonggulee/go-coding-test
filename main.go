package main

import (
	"fmt"
	"sort"
)

func solution(arr []int, divisor int) []int {
	res := []int{}
	for _, v := range arr {
		if v%divisor == 0 {
			res = append(res, v)
		}
	}

	if len(res) == 0 {
		return []int{-1}
	}

	sort.Ints(res)
	return res
}

func main() {
	r := solution([]int{5, 9, 7, 10}, 5)
	fmt.Println(r)

	r = solution([]int{2, 36, 1, 3}, 1)
	fmt.Println(r)

	r = solution([]int{3, 2, 6}, 10)
	fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
