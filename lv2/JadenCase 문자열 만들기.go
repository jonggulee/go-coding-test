import "strings"

func solution(s string) string {
	r := ""
	isNewWord := true

	for _, v := range s {
		if v == ' ' {
			r += " "
			isNewWord = true
			continue
		}

		if isNewWord {
			if v >= '0' && v <= '9' {
				r += string(v)
			} else if v >= 'a' && v <= 'z' {
				r += string(v - 32)
			} else if v >= 'A' && v <= 'Z' {
				r += string(v)
			}
			isNewWord = false
		} else {
			if v >= 'A' && v <= 'Z' {
				r += string(v + 32)
			} else {
				r += string(v)
			}
		}
	}

	return r
}
