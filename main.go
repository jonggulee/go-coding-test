package main

import (
	"fmt"
	"math"
)

func solution(a int, b int) int {
	if a%2 == 1 && b%2 == 1 {
		return (a * a) + (b * b)
	} else if a%2 == 1 && b%2 == 0 || a%2 == 0 && b%2 == 1 {
		return 2 * (a + b)
	} else {
		return int(math.Abs(float64(a - b)))
	}
}

func main() {
	r := solution(3, 5)
	fmt.Println(r)

	r = solution(6, 1)
	fmt.Println(r)

	r = solution(2, 4)
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
