package main

import (
	"fmt"
	"sort"
)

func solution(s string) string {
	bytes := []byte(s)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})

	return string(bytes)
}

func main() {
	r := solution("Zbcdefg")
	fmt.Println(r)

	// r = solution(24, 27)
	// fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
