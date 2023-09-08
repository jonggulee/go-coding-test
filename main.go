package main

import (
	"fmt"
	"math"
)

func solution(brown int, yellow int) []int {
	total := brown + yellow

	for height := 1; height <= int(math.Sqrt(float64(total))); height++ {
		width := total / height

		if width >= height && (width-2)*(height-2) == yellow {
			return []int{width, height}
		}
	}

	return []int{}
}

func main() {
	r := solution(10, 2)
	fmt.Println(r)

	r = solution(8, 1)
	fmt.Println(r)

	r = solution(24, 24)
	fmt.Println(r)
}
