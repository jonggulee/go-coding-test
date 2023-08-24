package main

import (
	"fmt"
)

func solution(arr []int) []int {
	stk := []int{}
	i := 0
	for i < len(arr) {
		if len(stk) == 0 {
			stk = append(stk, arr[i])
			i++
		} else if stk[len(stk)-1] < arr[i] {
			stk = append(stk, arr[i])
			i++
		} else {
			stk = stk[:len(stk)-1]
		}
	}
	return stk
}

func main() {
	r := solution([]int{1, 4, 2, 5, 3})
	fmt.Println(r)

	// r = solution("10203", "15")
	// fmt.Println(r)
}
