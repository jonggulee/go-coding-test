package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(binomial string) int {
	splitString := strings.Split(binomial, " ")
	a, _ := strconv.Atoi(splitString[0])
	operation := splitString[1]
	b, _ := strconv.Atoi(splitString[2])

	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	}

	return result
}

func main() {
	r := solution("43 + 12")
	fmt.Println(r)

	r = solution("0 - 7777")
	fmt.Println(r)

	r = solution("40000 * 40000")
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
