package main

import (
	"fmt"
)

func solution(arr []int, k int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range arr {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}

		if len(res) == k {
			break
		}
	}

	for len(res) < k {
		res = append(res, -1)
	}

	return res
}

func main() {
	r := solution([]int{0, 1, 1, 2, 2, 3}, 3)
	fmt.Println(r)

	r = solution([]int{0, 1, 1, 1, 1}, 4)
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
