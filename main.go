package main

import (
	"fmt"
)

func solution(n_str string) string {
	for {
		if n_str[0] == '0' {
			n_str = n_str[1:]
		} else {
			break
		}
	}

	return n_str
}

func main() {
	r := solution("0010")
	fmt.Println(r)

	r = solution("854020")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
