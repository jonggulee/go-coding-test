package main

import (
	"fmt"
	"sort"
)

func solution(num_list []int) []int {
	res := make([]int, 5)

	sort.Ints(num_list)

	for i := 0; i < 5; i++ {
		res[i] = num_list[i]
	}

	return res
}

func main() {
	r := solution([]int{12, 4, 15, 46, 38, 1, 14})
	fmt.Println(r)

	// r = solution([]int{444, 555, 666, 777}, 100)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
