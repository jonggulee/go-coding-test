import "strings"

func solution(myString string, pat string) int {
	chString := []rune{}
	for _, v := range myString {
		if v == 'A' {
			chString = append(chString, 'B')
		} else {
			chString = append(chString, 'A')
		}
	}

	if strings.Contains(string(chString), pat) {
		return 1
	}

	return 0
}