package main

import (
	"fmt"
)

func solution(n int, slicer []int, num_list []int) []int {
	switch n {
	case 1:
		return num_list[:slicer[1]+1]
	case 2:
		return num_list[slicer[0]:]
	case 3:
		return num_list[slicer[0] : slicer[1]+1]
	case 4:
		ret := []int{}
		for i := slicer[0]; i <= slicer[1]; i = i + slicer[2] {
			ret = append(ret, num_list[i])
		}
		return ret
	}
	return []int{}
}

func main() {
	r := solution(3, []int{1, 5, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(r)

	r = solution(4, []int{1, 5, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(r)

	// r = solution(24, 24)
	// fmt.Println(r)
}
