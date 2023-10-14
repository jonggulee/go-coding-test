package main

import (
	"fmt"
)

func solution(arr []int, flag []bool) []int {
	ret := []int{}
	for i, v := range flag {
		if v {
			for j := 0; j < arr[i]*2; j++ {
				ret = append(ret, arr[i])
			}
		} else {
			ret = ret[:len(ret)-arr[i]]
		}
	}
	return ret
}

func main() {
	r := solution([]int{3, 2, 4, 1, 3}, []bool{true, false, true, false, false})
	fmt.Println(r)

	// r = solution(2573)
	// fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
