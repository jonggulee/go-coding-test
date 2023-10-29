func solution(arr []int, delete_list []int) []int {
	res := []int{}
	deleteMap := make(map[int]bool)

	for _, d := range delete_list {
		deleteMap[d] = true
	}

	for _, v := range arr {
		if !deleteMap[v] {
			res = append(res, v)
		}
	}
	return res
}