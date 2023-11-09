package main

import (
	"fmt"
)

func solution(absolutes []int, signs []bool) int {
	var sum int
	for i, v := range absolutes {
		if signs[i] {
			sum += v
		} else {
			sum -= v
		}
	}
	return sum
}
func main() {
	r := solution([]int{4, 7, 12}, []bool{true, false, true})
	fmt.Println(r)

	r = solution([]int{1, 2, 3}, []bool{false, false, true})
	fmt.Println(r)

	// r = solution(626331)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
