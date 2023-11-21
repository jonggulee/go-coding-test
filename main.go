package main

import (
	"fmt"
	"strings"
)

func solution(s string) string {
	var ret string
	var indexCount int

	for _, v := range s {
		if v == ' ' {
			indexCount = 0
			ret += " "
			continue
		}
		if indexCount%2 == 0 {
			ret += strings.ToUpper(string(v))
		} else {
			ret += strings.ToLower(string(v))
		}
		indexCount++
	}
	return ret
}

func main() {
	r := solution("try hello world")
	fmt.Println(r)

	// r = solution([][]int{{1}, {2}}, [][]int{{3}, {4}})
	// fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
	// 116 - 84 = 32
	// 121 - 89 = 32

}
