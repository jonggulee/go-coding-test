package main

import (
	"fmt"
	"strconv"
)

func solution(n_str string) int {
	nInt, _ := strconv.Atoi(n_str)
	return nInt
}

func main() {
	r := solution("10")
	fmt.Println(r)

	r = solution("8542")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
