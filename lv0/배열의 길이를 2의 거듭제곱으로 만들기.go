func solution(arr []int) []int {
	length := len(arr)

	if length > 0 && (length&(length-1)) == 0 {
		return arr
	}

	next := 1
	for next < length {
		next *= 2
	}

	for i := length; i < next; i++ {
		arr = append(arr, 0)
	}

	return arr
}