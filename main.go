package main

import (
	"fmt"
	"sort"
	"strconv"
)

func solution(n int64) int64 {

	s := strconv.FormatInt(n, 10)

	runeSlice := []rune(s)

	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] > runeSlice[j]
	})

	sortedStr := string(runeSlice)

	ret, _ := strconv.ParseInt(sortedStr, 10, 64)

	return ret
}

func main() {
	r := solution(118372)
	fmt.Println(r)

	// r = solution("aaAA", "aaaaa")
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
