package main

import (
	"fmt"
)

func solution(n int) int {
	return n/7 + map[bool]int{true: 1, false: 0}[n%7 > 0]
}

func main() {
	r := solution(7)
	fmt.Println(r)

	r = solution(1)
	fmt.Println(r)

	r = solution(15)
	fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
	// 116 - 84 = 32
	// 121 - 89 = 32

}
