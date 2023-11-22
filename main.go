package main

import (
	"fmt"
)

func solution(number []int) int {
	count := 0

	for str := 0; str < len(number)-2; str++ {
		for mid := str + 1; mid < len(number)-1; mid++ {
			for end := mid + 1; end < len(number); end++ {
				if (number[str] + number[mid] + number[end]) == 0 {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	r := solution([]int{-2, 3, 0, 2, -5})
	fmt.Println(r)

	r = solution([]int{-3, -2, -1, 0, 1, 2, 3})
	fmt.Println(r)

	r = solution([]int{-1, 1, -1, 1})
	fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
	// 116 - 84 = 32
	// 121 - 89 = 32

}
