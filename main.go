package main

import (
	"fmt"
)

func solution(progresses []int, speeds []int) []int {
	for i := 0; i < len(progresses); i++ {
		if progresses[i] < 100 {
			progresses[i] += speeds[i]
		}
	}
	// return progresses
	return []int{}
}

func main() {
	r := solution([]int{93, 30, 55}, []int{1, 30, 5})
	fmt.Println(r)

	r = solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1})
	fmt.Println(r)

	// r = solution(-4, 2)
	// fmt.Println(r)
}
