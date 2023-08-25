import "strings"

func solution(my_string string, queries [][]int) string {
	var my_string_slice []string
	for _, v := range my_string {
		my_string_slice = append(my_string_slice, string(v))
	}

	for i := 0; i < len(queries); i++ {
		str := queries[i][0]
		end := queries[i][1]
		for j := str; j < end; j++ {
			my_string_slice[j], my_string_slice[end] = my_string_slice[end], my_string_slice[j]
			end--
		}
	}
	return strings.Join(my_string_slice, "")
}