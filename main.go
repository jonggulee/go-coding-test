package main

import (
	"fmt"
)

func solution(arr []int) []int {
	ret := []int{}
	for _, v := range arr {
		if v < 50 && v%2 == 1 {
			v = v * 2
			ret = append(ret, v)
		} else if v >= 50 && v%2 == 0 {
			v = v / 2
			ret = append(ret, v)
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	r := solution([]int{1, 2, 3, 100, 99, 98})
	fmt.Println(r)

	// r = solution([]int{58, 44, 27, 10, 100}, 139)
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
