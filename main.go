package main

import (
	"fmt"
)

func solution(arr []int) []int {
	length := len(arr)

	if length > 0 && (length&(length-1)) == 0 {
		return arr
	}

	next := 1
	for next < length {
		next *= 2
	}

	for i := length; i < next; i++ {
		arr = append(arr, 0)
	}

	return arr
}

func main() {
	r := solution([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(r)

	r = solution([]int{58, 172, 746, 89})
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
