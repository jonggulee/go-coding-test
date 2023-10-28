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