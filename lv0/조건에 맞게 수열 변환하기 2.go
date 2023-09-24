func solution(arr []int) int {
	count := 0
	for {
		changed := false
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] >= 50 && arr[i]%2 == 0 {
				changed = true
				arr[i] = arr[i] / 2
			} else if arr[i] < 50 && arr[i]%2 == 1 {
				arr[i] = arr[i]*2 + 1
				changed = true
			}
		}

		if !changed {
			return count
		}
		count++
	}
}