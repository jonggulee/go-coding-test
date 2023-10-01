func solution(strArr []string) []string {
	ret := []string{}
	for i := 0; i < len(strArr); i++ {
		if !strings.Contains(strArr[i], "ad") {
			ret = append(ret, strArr[i])
		}
	}
	return ret
}