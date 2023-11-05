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