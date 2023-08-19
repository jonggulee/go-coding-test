func solution(my_string string, k int) string {
	var r string
	for i := 0; i < k; i++ {
		r = r + my_string
	}
	return r
}