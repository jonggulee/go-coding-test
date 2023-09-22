package main

import (
	"fmt"
)

func solution(numbers []int, n int) int {
	ret := 0
	for i := 0; i < len(numbers); i++ {
		ret += numbers[i]
		if n < ret {
			return ret
		}
	}

	return 0
}

func main() {
	r := solution([]int{34, 5, 71, 29, 100, 34}, 123)
	fmt.Println(r)

	r = solution([]int{58, 44, 27, 10, 100}, 139)
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
