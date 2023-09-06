package main

import (
	"fmt"
)

func solution(q int, r int, code string) string {
	ret := ""
	for i := 0; i < len(code); i++ {
		if i%q == r {
			ret += string(code[i])
		}
	}
	return ret
}

func main() {
	r := solution(3, 1, "qjnwezgrpirldywt")
	fmt.Println(r)

	r = solution(1, 0, "programmers")
	fmt.Println(r)

	// r = solution(-4, 2)
	// fmt.Println(r)
}
