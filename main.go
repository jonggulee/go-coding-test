package main

import (
	"fmt"
)

func solution(order []string) int {
	sum := 0
	for _, v := range order {
		switch v {
		case "americano", "americanoice", "iceamericano", "anything":
			sum += 4500
		case "hotamericano", "americanohot":
			sum += 4500
		case "cafelatteice", "icecafelatte", "cafelatte":
			sum += 5000
		case "hotcafelatte", "cafelattehot":
			sum += 5000
		}
	}
	return sum
}

func main() {
	r := solution([]string{"cafelatte", "americanoice", "hotcafelatte", "anything"})
	fmt.Println(r)

	r = solution([]string{"americanoice", "americano", "iceamericano"})
	fmt.Println(r)

	// r = solution(2, 4)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
