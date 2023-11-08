package main

import (
	"fmt"
)

func solution(angle int) int {
	if angle > 0 && angle < 90 {
		return 1
	} else if angle == 90 {
		return 2
	} else if angle > 90 && angle < 180 {
		return 3
	} else if angle == 180 {
		return 4
	}
	return 0
}

func main() {
	r := solution(29183, 1)
	fmt.Println(r)

	r = solution(232443, 4)
	fmt.Println(r)

	r = solution(123456, 7)
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
