func solution(arr1 [][]int, arr2 [][]int) [][]int {
	res := make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = make([]int, len(arr1[i]))
		for j := 0; j < len(arr1[i]); j++ {
			res[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return res
}