package main

import (
	"fmt"
)

func solution(arr []int) []int {
	stk := []int{}
	for i := 0; i < len(arr); i++ {
		if len(stk) == 0 {
			stk = append(stk, arr[i])
		} else if stk[len(stk)-1] == arr[i] {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, arr[i])
		}
	}

	if len(stk) == 0 {
		return []int{-1}
	}

	return stk
}

func main() {
	r := solution([]int{0, 1, 1, 1, 0})
	fmt.Println(r)

	// r = solution(2573)
	// fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
