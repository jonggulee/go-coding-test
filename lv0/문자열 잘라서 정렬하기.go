func solution(myString string) []string {
	ret := []string{}
	splitX := strings.Split(myString, "x")
	for i := 0; i < len(splitX); i++ {
		if splitX[i] != "" {
			ret = append(ret, splitX[i])
		}
	}

	sort.Strings(ret)

	return ret
}