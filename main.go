package main

import (
	"fmt"
)

func solution(t string, p string) int {
	var r int
	for i := 0; i <= len(t)-len(p); i++ {
		// intVal := string(t[i : i+len(p)])
		if string(t[i:i+len(p)]) <= p {
			r += 1
		}

	}
	return r
}

func main() {
	r := solution("3141592", "271")
	fmt.Println(r)

	r = solution("500220839878", "7")
	fmt.Println(r)

	r = solution("10203", "15")
	fmt.Println(r)
}
