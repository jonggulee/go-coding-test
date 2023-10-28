package main

import (
	"fmt"
)

func solution(flo float64) int {
	return int(flo)
}

func main() {
	r := solution(1.42)
	fmt.Println(r)

	// r = solution([]int{444, 555, 666, 777}, 100)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
