package main

import (
	"fmt"
)

func solution(arr1 []int, arr2 []int) int {
	sumArr1 := 0
	sumArr2 := 0

	if len(arr1) > len(arr2) {
		return 1
	} else if len(arr1) < len(arr2) {
		return -1
	} else {
		for i := 0; i < len(arr1); i++ {
			sumArr1 += arr1[i]
			sumArr2 += arr2[i]
		}

		if sumArr1 == sumArr2 {
			return 0
		} else if sumArr1 > sumArr2 {
			return 1
		} else {
			return -1
		}
	}
}

func main() {
	r := solution([]int{49, 13}, []int{70, 11, 2})
	fmt.Println(r)

	r = solution([]int{100, 17, 84, 1}, []int{55, 12, 65, 36})
	fmt.Println(r)

	r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
