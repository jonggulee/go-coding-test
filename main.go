package main

import (
	"fmt"
	"strings"
)

func solution(str1 string, str2 string) int {
	if strings.Contains(str2, str1) {
		return 1
	}
	return 0
}

func main() {
	r := solution("abc", "aabcc")
	fmt.Println(r)

	r = solution("tbt", "tbbttb")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
