func solution(myString string) string {
	byteArray := []byte(myString)
	for i := 0; i < len(myString); i++ {
		if byteArray[i] < 'l' {
			byteArray[i] = 'l'
		}
	}
	return string(byteArray)
}
