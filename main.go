package main

import (
	"fmt"
)

type Item struct {
	priority int
	location int
}

func solution(priorities []int, location int) int {
	queue := make([]Item, len(priorities))
	for i, q := range priorities {
		queue[i] = Item{q, i}
	}

	count := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		highest := true
		for _, q := range queue {
			if q.priority > item.priority {
				highest = false
				break
			}
		}

		if highest {
			count++
			if item.location == location {
				return count
			}
		} else {
			queue = append(queue, item)
		}
	}

	return 0
}

func main() {
	r := solution([]int{2, 1, 3, 2}, 2)
	fmt.Println(r)

	r = solution([]int{1, 1, 9, 1, 1, 1}, 0)
	fmt.Println(r)

	// r = solution([]int{1, 2, 3, 4, 5}, []int{3, 3, 3, 3, 3})
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
