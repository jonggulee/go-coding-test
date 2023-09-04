package main

import (
	"fmt"
)

func solution(my_string string, s int, e int) string {
	myBytes := []byte(my_string)

	for i, j := s, e; i < j; i, j = i+1, j-1 {
		myBytes[i], myBytes[j] = myBytes[j], myBytes[i]
	}

	return string(myBytes)

	// return ""
}

func main() {
	r := solution("Progra21Sremm3", 6, 12)
	fmt.Println(r)

	r = solution("Stanley1yelnatS", 4, 10)
	fmt.Println(r)

	// r = solution(-4, 2)
	// fmt.Println(r)
}
