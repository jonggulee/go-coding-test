package main

import (
	"fmt"
)

func solution(str_list []string) []string {
	for i, v := range str_list {
		if v == "l" {
			return str_list[:i]
		} else if v == "r" {
			return str_list[i+1:]
		}
	}
	return []string{}
}

func main() {
	r := solution([]string{"u", "u", "l", "r"})
	fmt.Println(r)

	r = solution([]string{"l"})
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
