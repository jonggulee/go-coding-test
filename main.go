package main

import (
	"fmt"
	"strconv"
)

func solution(l int, r int) []int {
	var ret []int
	for num := l; num <= r; num++ {
		strNum := strconv.Itoa(num)
		allValid := true

		for _, digit := range strNum {
			if digit != '0' && digit != '5' {
				fmt.Println(digit, strNum)
				allValid = false
				break
			}
		}

		if allValid {
			ret = append(ret, num)
		}
	}

	if len(ret) == 0 {
		return []int{-1}
	}

	return ret
}

func main() {
	r := solution(5, 555)
	fmt.Println(r)

	r = solution(10, 20)
	fmt.Println(r)

	// r = solution("10203", "15")
	// fmt.Println(r)
}
