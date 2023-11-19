package main

import (
	"fmt"
)

func solution(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}

	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func main() {
	r := solution("a234")
	fmt.Println(r)

	r = solution("1234")
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
