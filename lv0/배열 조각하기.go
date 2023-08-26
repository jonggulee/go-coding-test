func solution(arr []int, query []int) []int {
	for i, q := range query {
		if i%2 == 0 {
			arr = arr[:q+1]
		} else {
			arr = arr[q:]
		}
	}
	return arr
}