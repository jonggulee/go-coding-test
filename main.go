package main

import (
	"fmt"
	"strconv"
)

func solution(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}

func main() {
	r := solution("-1234")
	fmt.Println(r)

	// r = solution("aaAA", "aaaaa")
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
