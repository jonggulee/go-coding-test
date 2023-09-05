package main

import (
	"fmt"
)

func solution(my_string string, m int, c int) string {
	r := ""
	for i := 0; i < len(my_string); i = i + m {
		r = r + string(my_string[i : m+i][c-1])
	}
	return r
}

func main() {
	r := solution("ihrhbakrfpndopljhygc", 4, 2)
	fmt.Println(r)

	r = solution("programmers", 1, 1)
	fmt.Println(r)

	// r = solution(-4, 2)
	// fmt.Println(r)
}
