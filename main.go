package main

import (
	"fmt"
	"strings"
)

func solution(myString string, pat string) int {
	lowString := strings.ToLower(myString)
	lowPat := strings.ToLower(pat)

	if strings.Contains(lowString, lowPat) {
		return 1
	}

	return 0
}

func main() {
	r := solution("AbCdEfG", "aBc")
	fmt.Println(r)

	r = solution("aaAA", "aaaaa")
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
