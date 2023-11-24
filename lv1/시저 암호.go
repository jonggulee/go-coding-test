func solution(s string, n int) string {
	ret := ""

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			ret += string((v-'a'+rune(n))%26 + 'a')
		} else if v >= 'A' && v <= 'Z' {
			ret += string((v-'A'+rune(n))%26 + 'A')
		} else {
			ret += " "
		}
	}
	return ret
}