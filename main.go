package main

import (
	"fmt"
)

func solution(myString string) string {
	byteArray := []byte(myString)
	for i := 0; i < len(myString); i++ {
		if byteArray[i] < 'l' {
			byteArray[i] = 'l'
		}
	}
	return string(byteArray)
}

func main() {
	r := solution("abcdevwxyz")
	fmt.Println(r)

	r = solution("jjnnllkkmm")
	fmt.Println(r)

	// r = solution(2, 4)
	// fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
