func solution(my_string string, s int, e int) string {
	myBytes := []byte(my_string)

	for i, j := s, e; i < j; i, j = i+1, j-1 {
		myBytes[i], myBytes[j] = myBytes[j], myBytes[i]
	}

	return string(myBytes)
}