package main

import (
	"fmt"
)

func solution(str1 string, str2 string) string {
	var r string
	for i := 0; i < len(str1); i++ {
		r = r + string(str1[i]) + string(str2[i])
	}
	return r
}

func main() {
	r := solution("aaaaa", "bbbbb")
	fmt.Println(r)

}
