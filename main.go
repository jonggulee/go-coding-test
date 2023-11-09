package main

import (
	"fmt"
)

func solution(x int) bool {
	original := x
	sumOfDigits := 0

	for x > 0 {
		fmt.Println(x, sumOfDigits)
		sumOfDigits += x % 10
		x /= 10
	}

	return original%sumOfDigits == 0
}

func main() {
	r := solution(10)
	fmt.Println(r)

	r = solution(12)
	fmt.Println(r)

	r = solution(11)
	fmt.Println(r)

	r = solution(13)
	fmt.Println(r)
}
