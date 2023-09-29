import (
	"unicode"
	"strings"
)

func solution(myString string) string {
	var builder strings.Builder

	for _, v := range myString {
		switch {
		case v == 'a' || v == 'A':
			builder.WriteRune('A')
		case unicode.IsUpper(v):
			builder.WriteRune(unicode.ToLower(v))
		default:
			builder.WriteRune(v)
		}
	}
	return builder.String()
}