import "sort"

func solution(my_string string) []string {
	r := []string{}
	for i := 0; i < len(my_string); i++ {
		r = append(r, string(my_string[i:]))
	}
	sort.Strings(r)

	return r
}