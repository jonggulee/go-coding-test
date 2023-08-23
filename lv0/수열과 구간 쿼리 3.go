func solution(arr []int, queries [][]int) []int {
	for i := 0; i < len(queries); i++ {
		str := queries[i][0]
		end := queries[i][1]
		arr[str], arr[end] = arr[end], arr[str]
	}
	return arr
}