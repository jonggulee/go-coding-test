package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(myString string) []string {
	ret := []string{}
	splitX := strings.Split(myString, "x")
	for i := 0; i < len(splitX); i++ {
		if splitX[i] != "" {
			ret = append(ret, splitX[i])
		}
	}

	sort.Strings(ret)

	return ret
}

func main() {
	r := solution("axbxcxdx")
	fmt.Println(r)

	r = solution("dxccxbbbxaaaa")
	fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
