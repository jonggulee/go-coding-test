func solution(names []string) []string {
	ret := []string{}
	for i := 0; i < len(names); i += 5 {
		ret = append(ret, names[i])
	}
	return ret
}