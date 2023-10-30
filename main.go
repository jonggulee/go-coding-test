package main

import (
	"fmt"
	"strings"
)

func solution(str_list []string, ex string) string {
	res := ""

	for _, str := range str_list {
		if !strings.Contains(str, ex) {
			res += str
		}
	}

	return res
}

func main() {
	r := solution([]string{"abc", "def", "ghi"}, "ef")
	fmt.Println(r)

	r = solution([]string{"abc", "bbc", "cbc"}, "c")
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
