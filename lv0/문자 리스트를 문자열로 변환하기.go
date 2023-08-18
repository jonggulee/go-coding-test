func solution(arr []string) string {
	var r string
	for i := 0; i < len(arr); i++ {
		r = r + arr[i]
	}
	return r
}