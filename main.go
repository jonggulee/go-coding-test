package main

import (
	"fmt"
)

func solution(a int, d int, included []bool) int {
	var seq = a
	var r int
	for i := 0; i < len(included); i++ {
		if included[i] {
			r = r + seq
		}
		seq = seq + d
	}
	return r
}

func main() {
	r := solution(3, 4, []bool{true, false, false, true, true})
	fmt.Println(r)

	r = solution(7, 1, []bool{false, false, false, true, false, false, false})
	fmt.Println(r)

}
