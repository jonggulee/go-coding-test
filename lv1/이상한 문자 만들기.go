import "strings"

func solution(s string) string {
	var ret string
	var indexCount int

	for _, v := range s {
		if v == ' ' {
			indexCount = 0
			ret += " "
			continue
		}
		if indexCount%2 == 0 {
			ret += strings.ToUpper(string(v))
		} else {
			ret += strings.ToLower(string(v))
		}
		indexCount++
	}
	return ret
}