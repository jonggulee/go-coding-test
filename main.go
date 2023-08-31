package main

import (
	"fmt"
)

func solution(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ret += n / i
		}
	}
	return ret
}

func main() {
	r := solution(12)
	fmt.Println(r)

	r = solution(5)
	fmt.Println(r)

	// r = solution("-1 -1")
	// fmt.Println(r)
}
