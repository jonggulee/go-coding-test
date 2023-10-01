package main

import (
	"fmt"
	"strings"
)

func solution(my_string string) []string {
	return strings.Split(my_string, " ")
}

func main() {
	r := solution("i love you")
	fmt.Println(r)

	// r = solution("programmers")
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
