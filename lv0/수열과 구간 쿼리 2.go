func solution(arr []int, queries [][]int) []int {
	results := make([]int, len(queries))

	for i, query := range queries {
		start := query[0]
		end := query[1]
		k := query[2]
		minValue := math.MaxInt32

		for j := start; j <= end; j++ {
			if arr[j] > k && arr[j] < minValue {
				minValue = arr[j]
			}
		}

		if minValue == math.MaxInt32 {
			results[i] = -1
		} else {
			results[i] = minValue
		}
	}

	return results
}