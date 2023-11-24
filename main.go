package main

import (
	"fmt"
)

func solution(s string, n int) string {
	ret := ""

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			ret += string((v-'a'+rune(n))%26 + 'a')
		} else if v >= 'A' && v <= 'Z' {
			ret += string((v-'A'+rune(n))%26 + 'A')
		} else {
			ret += " "
		}

		// if v == ' ' {
		// 	ret += string(' ')
		// } else if nextStr <= 'z' && nextStr >= 'a' || nextStr <= 'Z' && nextStr >= 'A' {
		// 	ret += string(v + rune(n))
		// } else if nextStr >= 'z' {
		// 	ret += string('a' + nextStr - 'z' - 1)
		// } else if nextStr >= 'Z' {
		// 	ret += string('A' + nextStr - 'Z' - 1)
		// }
	}
	return ret
}

func main() {
	r := solution("AB", 1)
	fmt.Println(r)

	r = solution("z", 3)
	fmt.Println(r)

	r = solution("a B z", 4)
	fmt.Println(r)

	// r = solution(13)
	// fmt.Println(r)
	// 116 - 84 = 32
	// 121 - 89 = 32

}
