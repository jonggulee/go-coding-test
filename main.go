package main

import (
	"fmt"
)

func solution(arr []int) int {
	count := 0

	for {
		changed := false
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] >= 50 && arr[i]%2 == 0 {
				changed = true
				arr[i] = arr[i] / 2
			} else if arr[i] < 50 && arr[i]%2 == 1 {
				arr[i] = arr[i]*2 + 1
				changed = true
			}
		}

		if !changed {
			return count
		}
		count++
	}
}

func main() {
	r := solution([]int{1, 2, 3, 100, 99, 98})
	fmt.Println(r)

	// r = solution([]int{58, 44, 27, 10, 100}, 139)
	// fmt.Println(r)

	// r = solution([]int{1, 1, 1})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
