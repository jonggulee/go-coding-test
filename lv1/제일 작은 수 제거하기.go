func solution(arr []int) []int {
	res := []int{}
	if len(arr) <= 1 {
		return []int{-1}
	}

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != min {
			res = append(res, arr[i])
		}
	}

	return res
}