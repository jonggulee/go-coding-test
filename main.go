package main

import (
	"fmt"
)

func solution(arr [][]int) [][]int {
	rowLength := len(arr)
	colLength := len(arr[0])

	if rowLength > colLength {
		for i := range arr {
			for len(arr[i]) < rowLength {
				arr[i] = append(arr[i], 0)
			}
		}
	}

	if colLength > rowLength {
		for rowLength < colLength {
			newRow := make([]int, colLength)
			arr = append(arr, newRow)
			rowLength++
		}
	}

	return arr
}

func main() {
	r := solution([][]int{{572, 22, 37}, {287, 726, 384}, {85, 137, 292}, {487, 13, 876}})
	fmt.Println(r)

	r = solution([][]int{{57, 192, 534, 2}, {9, 345, 192, 999}})
	fmt.Println(r)

	r = solution([][]int{{1, 2}, {3, 4}})
	fmt.Println(r)

	// r = solution([]int{1, 2, 1, 2, 1, 10, 2, 1})
	// fmt.Println(r)
}
