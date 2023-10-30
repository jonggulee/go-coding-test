import "strings"

func solution(str_list []string, ex string) string {
	res := ""

	for _, str := range str_list {
		if !strings.Contains(str, ex) {
			res += str
		}
	}

	return res
}