import "strings"

func solution(strArr []string) []string {
	for i := 0; i < len(strArr); i++ {
		if i%2 == 0 {
			strArr[i] = strings.ToLower(strArr[i])
		} else {
			strArr[i] = strings.ToUpper(strArr[i])
		}
	}
	return strArr
}