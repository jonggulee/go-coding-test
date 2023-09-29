package main

import (
	"fmt"
	"strings"
)

func solution(myString string, pat string) string {
	lastIndex := strings.LastIndex(myString, pat)
	return myString[:lastIndex+len(pat)]
}

func main() {
	r := solution("AbCdEFG", "dE")
	fmt.Println(r)

	r = solution("AAAAaaaa", "a")
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
