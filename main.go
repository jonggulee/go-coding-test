package main

import (
	"fmt"
	"strings"
)

func solution(myString string) string {
	return strings.ToLower(myString)
}

func main() {
	r := solution("aBcDeFg")
	fmt.Println(r)

	// r = solution("aaAA", "aaaaa")
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
