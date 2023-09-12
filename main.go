package main

import (
	"fmt"
)

func solution(arr []int) []int {
	firstIndex := -1
	lastIndex := -1

	for i, v := range arr {
		if v == 2 {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	if lastIndex == -1 {
		return []int{-1}
	}

	return arr[firstIndex : lastIndex+1]
}

func main() {
	r := solution([]int{1, 2, 1, 4, 5, 2, 9})
	fmt.Println(r)

	r = solution([]int{1, 2, 1})
	fmt.Println(r)

	r = solution([]int{1, 1, 1})
	fmt.Println(r)

	r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	fmt.Println(r)
}
