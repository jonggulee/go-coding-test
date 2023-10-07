package main

import (
	"fmt"
)

func solution(progresses []int, speeds []int) []int {
	var result []int
	var completeDays []int

	for i := 0; i < len(progresses); i++ {
		days := (100 - progresses[i] + speeds[i] - 1) / speeds[i]
		completeDays = append(completeDays, days)
	}

	for len(completeDays) > 0 {
		first := completeDays[0]
		count := 0

		for len(completeDays) > 0 && completeDays[0] <= first {
			count++
			completeDays = completeDays[1:]
		}

		result = append(result, count)
	}

	return result
}

func main() {
	r := solution([]int{93, 30, 55}, []int{1, 30, 5})
	fmt.Println(r)

	r = solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1})
	fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
