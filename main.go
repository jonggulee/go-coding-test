package main

import (
	"fmt"
	"strings"
)

func solution(strArr []string) []string {
	ret := []string{}
	for i := 0; i < len(strArr); i++ {
		if !strings.Contains(strArr[i], "ad") {
			ret = append(ret, strArr[i])
		}
	}
	return ret
}

func main() {
	r := solution([]string{"and", "notad", "abcd"})
	fmt.Println(r)

	// r = solution([]string{"there", "are", "no", "a", "ds"})
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
