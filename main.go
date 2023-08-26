package main

import (
	"fmt"
	"sort"
)

func solution(rank []int, attendance []bool) int {
	rankMap := make(map[int]bool)
	rankIndex := make(map[int]int)
	top := []int{}
	for i, v := range rank {
		rankMap[v] = attendance[i]
		rankIndex[v] = i
	}

	for k, v := range rankMap {
		if v {
			top = append(top, k)
		}
	}
	sort.Ints(top)

	return (10000 * rankIndex[top[0]]) + (100 * rankIndex[top[1]]) + rankIndex[top[2]]
}

func main() {
	r := solution([]int{3, 7, 2, 5, 4, 6, 1}, []bool{false, true, true, true, true, false, false})
	fmt.Println(r)

	r = solution([]int{1, 2, 3}, []bool{true, true, true})
	fmt.Println(r)

	r = solution([]int{6, 1, 5, 2, 3, 4}, []bool{true, false, true, false, false, true})
	fmt.Println(r)
}
