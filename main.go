package main

import (
	"fmt"
)

func solution(num_str string) int {
	res := 0
	for _, v := range num_str {
		res += int(v - '0')
	}

	return res
}

func main() {
	r := solution("123456789")
	fmt.Println(r)

	r = solution("1000000")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
