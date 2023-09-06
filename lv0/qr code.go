func solution(q int, r int, code string) string {
	ret := ""
	for i := 0; i < len(code); i++ {
		if i%q == r {
			ret += string(code[i])
		}
	}
	return ret
}