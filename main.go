package main

import (
	"fmt"
)

func solution(price int, money int, count int) int64 {
	sum := 0

	for i := 1; i <= count; i++ {
		sum += (price * i)
	}

	if sum < money {
		return 0
	} else {
		return int64(sum - money)
	}
}

func main() {
	r := solution(3, 20, 4)
	fmt.Println(r)

	// r = solution(24, 27)
	// fmt.Println(r)

	// r = solution([]int{3, 2, 6}, 10)
	// fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
}
