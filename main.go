package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution(s string) string {
	arrStr := strings.Split(s, " ")
	arrInt := []int{}

	for _, str := range arrStr {
		num, _ := strconv.Atoi(str)
		arrInt = append(arrInt, num)
	}

	sort.Ints(arrInt)

	min := strconv.Itoa(arrInt[0])
	max := strconv.Itoa(arrInt[len(arrInt)-1])
	return min + " " + max
}

func main() {
	r := solution("1 2 3 4")
	fmt.Println(r)

	r = solution("-1 -2 -3 -4")
	fmt.Println(r)

	r = solution("-1 -1")
	fmt.Println(r)
}
