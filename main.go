package main

import (
	"fmt"
)

func solution(x int, n int) (answer []int64) {
	for i := 1; i <= n; i++ {
		answer = append(answer, int64(i*x))
	}
	return
}

func main() {
	r := solution(2, 5)
	fmt.Println(r)

	r = solution(4, 3)
	fmt.Println(r)

	r = solution(-4, 2)
	fmt.Println(r)
}
