package main

import (
	"fmt"
	"strings"
)

func solution(my_string string, target string) int {
	if strings.Contains(my_string, target) {
		return 1
	}
	return 0
}

func main() {
	r := solution("banana", "ana")
	fmt.Println(r)

	r = solution("banana", "wxyz")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
