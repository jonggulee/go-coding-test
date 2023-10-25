func solution(arr []int, n int) []int {
	for i := 0; len(arr) > i; i++ {
		if len(arr)%2 != 0 {
			if i%2 == 0 {
				arr[i] = arr[i] + n
			}
		} else {
			if i%2 != 0 {
				arr[i] = arr[i] + n
			}
		}
	}
	return arr
}