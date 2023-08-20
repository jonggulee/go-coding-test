func solution(code string) string {
	var ret string
	var mode = "0"
	for i := 0; i < len(code); i++ {
		if mode == "0" {
			if string(code[i]) == "1" {
				mode = "1"
			} else if i%2 == 0 {
				ret += string(code[i])
			}
		} else if mode == "1" {
			if string(code[i]) == "1" {
				mode = "0"
			} else if i%2 == 1 {
				ret += string(code[i])
			}
		}
	}
	if ret == "" {
		return "EMPTY"
	}
	return ret
}