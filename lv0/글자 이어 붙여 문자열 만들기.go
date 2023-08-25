import "strings"

func solution(my_string string, index_list []int) string {
	var r []string
	for _, index := range index_list {
		r = append(r, string(my_string[index]))
	}
	return strings.Join(r, "")
}