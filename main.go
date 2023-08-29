package main

import (
	"fmt"
)

// func solution(s string) string {
// 	r := ""
// 	arrStr := strings.Split(s, " ")
// 	for _, str := range arrStr {
// 		for i, v := range str {
// 			if i == 0 {
// 				if v >= 48 && v <= 57 {
// 					r += string(v)
// 				} else if v >= 97 && v <= 122 {
// 					upper := v - 32
// 					r += string(upper)
// 				}
// 			} else {
// 				if v >= 65 && v <= 90 {
// 					lower := v + 32
// 					r += string(lower)
// 				} else {
// 					r += string(v)
// 				}
// 			}
// 		}
// 		if str != arrStr[len(arrStr)-1] {
// 			r += " "
// 		}
// 	}
// 	return r
// }

func solution(s string) string {
	r := ""
	isNewWord := true

	for _, v := range s {
		if v == ' ' {
			r += " "
			isNewWord = true
			continue
		}

		if isNewWord {
			if v >= '0' && v <= '9' {
				r += string(v)
			} else if v >= 'a' && v <= 'z' {
				r += string(v - 32)
			} else if v >= 'A' && v <= 'Z' {
				r += string(v)
			}
			isNewWord = false
		} else {
			if v >= 'A' && v <= 'Z' {
				r += string(v + 32)
			} else {
				r += string(v)
			}
		}
	}

	return r
}

func main() {
	r := solution("3people unFollowed me")
	fmt.Println(r)

	r = solution("for the last week")
	fmt.Println(r)

	// r = solution("-1 -1")
	// fmt.Println(r)
}
