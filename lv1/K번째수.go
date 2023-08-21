func solution(array []int, commands [][]int) []int {
	var r []int
	for i := 0; i < len(commands); i++ {
		arr := append([]int{}, array[commands[i][0]-1:commands[i][1]]...)

		for j := 0; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[j] > arr[k] {
					arr[j], arr[k] = arr[k], arr[j]
				}
			}
		}
		r = append(r, arr[commands[i][2]-1])
	}
	return r
}