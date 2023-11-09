package main

import (
	"fmt"
)

func solution(num int) int {
	if num == 1 {
		return 0
	}

	for i := 1; i <= 500; i++ {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}

		if num == 1 {
			return i
		}
	}

	return -1
}

func main() {
	r := solution(6)
	fmt.Println(r)

	r = solution(16)
	fmt.Println(r)

	r = solution(626331)
	fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
