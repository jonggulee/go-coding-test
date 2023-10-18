func solution(arr []int, k int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range arr {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}

		if len(res) == k {
			break
		}
	}

	for len(res) < k {
		res = append(res, -1)
	}

	return res
}