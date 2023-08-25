func solution(arr []int, queries [][]int) []int {
	for _, query := range queries {
		s, e, k := query[0], query[1], query[2]
		for i := s; i <= e; i++ {
			if i%k == 0 {
				arr[i] += 1
			}
		}
	}
	return arr
}