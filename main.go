package main

import (
	"fmt"
)

func solution(my_string string, overwrite_string string, s int) string {
	s_word := my_string[:s]
	overwirte_len := len(overwrite_string)
	l_word := my_string[len(s_word)+overwirte_len:]
	return s_word + overwrite_string + l_word
}

func main() {
	r := solution("He11oWor1d", "lloWorl", 2)
	fmt.Println(r)

	a := solution("Program29b8UYP", "merS123", 7)
	fmt.Println(a)
}
