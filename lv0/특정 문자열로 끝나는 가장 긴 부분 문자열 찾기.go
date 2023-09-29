import "strings"

func solution(myString string, pat string) string {
	lastIndex := strings.LastIndex(myString, pat)
	return myString[:lastIndex+len(pat)]
}