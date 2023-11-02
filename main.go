package main

import (
	"fmt"
)

func solution(picture []string, k int) []string {
	var res []string
	for _, row := range picture {
		expandedRow := ""
		for _, pixel := range row {
			for i := 0; i < k; i++ {
				expandedRow += string(pixel)
			}
		}

		for i := 0; i < k; i++ {
			res = append(res, expandedRow)
		}
	}
	return res
}

func main() {
	r := solution([]string{".xx...xx.", "x..x.x..x", "x...x...x", ".x.....x.", "..x...x..", "...x.x...", "....x...."}, 2)
	fmt.Println(r)

	r = solution([]string{"x.x", ".x.", "x.x"}, 3)
	fmt.Println(r)

	// r = solution(2, 4)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
