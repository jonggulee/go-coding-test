package main

import (
	"fmt"
	"strconv"
)

func solution(a string, b string) string {
	// 문자열 뒤집기
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	// 뒤집어진 문자열로 덧셈
	a, b = reverse(a), reverse(b)
	lenA, lenB := len(a), len(b)
	carry := 0
	result := ""

	for i := 0; i < lenA || i < lenB || carry > 0; i++ {
		sum := carry
		if i < lenA {
			sum += int(a[i] - '0')
		}
		if i < lenB {
			sum += int(b[i] - '0')
		}
		carry = sum / 10
		result += strconv.Itoa(sum % 10)
	}

	return reverse(result)
}

func main() {
	r := solution("582", "734")
	fmt.Println(r)

	// r = solution(2573)
	// fmt.Println(r)

	// r = solution("40000 * 40000")
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
