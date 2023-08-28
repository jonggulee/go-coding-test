package main

import (
	"fmt"
	"sort"
)

func solution(my_string string) []string {
	r := []string{}
	for i := 0; i < len(my_string); i++ {
		// fmt.Println(string(my_string[i:]))
		r = append(r, string(my_string[i:]))
	}
	sort.Strings(r)

	return r
}

func main() {
	r := solution("banana")
	fmt.Println(r)

	r = solution("programmers")
	fmt.Println(r)

	// r = solution([]int{6, 1, 5, 2, 3, 4}, []bool{true, false, true, false, false, true})
	// fmt.Println(r)
}
