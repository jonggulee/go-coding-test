func solution(my_strings []string, parts [][]int) string {
	var ret string
	for i, str := range my_strings {
		ret = ret + str[parts[i][0]:parts[i][1]+1]
	}
	return ret
}