func solution(arr []int, queries [][]int) []int {
	for _, query := range queries {
		s, e := query[0], query[1]
		for i := s; i <= e; i++ {
			arr[i]++
		}
	}
	return arr
}