package main

import (
	"fmt"
)

func solution(my_strings []string, parts [][]int) string {
	var ret string
	for i, str := range my_strings {
		ret = ret + str[parts[i][0]:parts[i][1]+1]
	}
	return ret
}

func main() {
	r := solution([]string{"progressive", "hamburger", "hammer", "ahocorasick"}, [][]int{{0, 4}, {1, 2}, {3, 5}, {7, 7}})
	fmt.Println(r)

	// r = solution([]int{1, 2, 3}, []bool{true, true, true})
	// fmt.Println(r)

	// r = solution([]int{6, 1, 5, 2, 3, 4}, []bool{true, false, true, false, false, true})
	// fmt.Println(r)
}
