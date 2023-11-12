package main

import (
	"fmt"
)

func solution(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		divCount := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divCount += 1
			}
		}

		if divCount%2 == 0 {
			res += i
		} else {
			res -= i
		}
	}

	return res
}

func main() {
	r := solution(13, 17)
	fmt.Println(r)

	r = solution(24, 27)
	fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
