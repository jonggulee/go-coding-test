package main

import (
	"fmt"
)

func solution(arr []int, n int) []int {
	for i := 0; len(arr) > i; i++ {
		if len(arr)%2 != 0 {
			if i%2 == 0 {
				arr[i] = arr[i] + n
			}
		} else {
			if i%2 != 0 {
				arr[i] = arr[i] + n
			}
		}
	}
	return arr
}

func main() {
	r := solution([]int{49, 12, 100, 276, 33}, 27)
	fmt.Println(r)

	r = solution([]int{444, 555, 666, 777}, 100)
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
