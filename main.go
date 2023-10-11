package main

import (
	"fmt"
)

func solution(arr []int) []int {
	var ret []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func main() {
	r := solution([]int{5, 1, 4})
	fmt.Println(r)

	r = solution([]int{6, 6})
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
