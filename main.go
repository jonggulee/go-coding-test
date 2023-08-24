package main

import (
	"fmt"
)

func solution(x1 bool, x2 bool, x3 bool, x4 bool) bool {
	return (x1 || x2) && (x3 || x4)
}

func main() {
	r := solution(false, true, true, true)
	fmt.Println(r)

	r = solution(true, false, false, false)
	fmt.Println(r)
}
