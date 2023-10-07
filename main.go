package main

import (
	"fmt"
	"strings"
)

func solution(myString string, pat string) int {
	chString := []rune{}
	for _, v := range myString {
		if v == 'A' {
			chString = append(chString, 'B')
		} else {
			chString = append(chString, 'A')
		}
	}

	if strings.Contains(string(chString), pat) {
		return 1
	}

	return 0
}

func main() {
	r := solution("ABBAA", "AABB")
	fmt.Println(r)

	r = solution("ABAB", "ABAB")
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
