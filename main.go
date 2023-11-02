package main

import (
	"fmt"
)

func solution(date1 []int, date2 []int) int {
	for i := 0; i < 3; i++ {
		if date1[i] < date2[i] {
			return 1
		} else if date1[i] > date2[i] {
			return 0
		}
	}
	return 0
}

func main() {
	r := solution([]int{2021, 12, 28}, []int{2021, 12, 29})
	fmt.Println(r)

	r = solution([]int{1024, 10, 24}, []int{1024, 10, 24})
	fmt.Println(r)

	// r = solution(2, 4)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
