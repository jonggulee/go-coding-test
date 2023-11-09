package main

import (
	"fmt"
)

func solution(numbers []int) int {
	res := 45
	for _, v := range numbers {
		res -= v
	}

	return res
}

func main() {
	r := solution([]int{1, 2, 3, 4, 6, 7, 8, 0})
	fmt.Println(r)

	r = solution([]int{5, 8, 4, 0, 6, 7, 9})
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
