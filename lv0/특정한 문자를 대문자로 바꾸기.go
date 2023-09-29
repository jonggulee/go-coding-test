import "strings"

func solution(my_string string, alp string) string {
	for i := 0; i < len(my_string); i++ {
		if string(my_string[i]) == alp {
			my_string = my_string[:i] + strings.ToUpper(string(my_string[i])) + my_string[i+1:]
		}
	}
	return my_string
}