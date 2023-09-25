package main

import (
	"fmt"
)

func solution(num_list []int) int {
	count := 0
	for i := 0; i < len(num_list); i++ {
		for num_list[i] != 1 {
			if num_list[i]%2 == 1 {
				num_list[i] = (num_list[i] - 1) / 2
				count++
			} else if num_list[i]%2 == 0 {
				num_list[i] /= 2
				count++
			}
		}
	}

	return count
}

func main() {
	r := solution([]int{12, 4, 15, 1, 14})
	fmt.Println(r)

	// r = solution([]int{58, 44, 27, 10, 100}, 139)
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
