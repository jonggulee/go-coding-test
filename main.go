package main

import (
	"fmt"
	"strings"
)

func solution(myString string) []int {
	ret := []int{}
	splitX := strings.Split(myString, "x")
	for i := 0; i < len(splitX); i++ {
		count := len(splitX[i])
		ret = append(ret, count)
	}
	return ret
}

func main() {
	r := solution("oxooxoxxox")
	fmt.Println(r)

	// r = solution("programmers")
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
