package main

import (
	"fmt"
	"math"
)

func solution(n int64) int64 {
	sqrt := math.Sqrt(float64(n))

	if sqrt == float64(int64(sqrt)) {
		return int64(sqrt+1) * int64(sqrt+1)
	}

	return -1
}

func main() {
	r := solution(121)
	fmt.Println(r)

	r = solution(12)
	fmt.Println(r)

	// r = solution(123456, 7)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
