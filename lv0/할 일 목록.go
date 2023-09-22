func solution(todo_list []string, finished []bool) []string {
	ret := []string{}
	for i := 0; i < len(todo_list); i++ {
		if finished[i] == false {
			ret = append(ret, todo_list[i])
		}
	}

	return ret
}