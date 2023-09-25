package main

import (
	"fmt"
)

func solution(num_list []int) int {
	if len(num_list) >= 11 {
		sum := 0
		for _, num := range num_list {
			sum += num
		}
		return sum
	} else {
		product := 1
		for _, num := range num_list {
			product *= num
		}
		return product
	}
}

func main() {
	r := solution([]int{3, 4, 5, 2, 5, 4, 6, 7, 3, 7, 2, 2, 1})
	fmt.Println(r)

	r = solution([]int{2, 3, 4, 5})
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
