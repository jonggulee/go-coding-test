func solution(binomial string) int {
	splitString := strings.Split(binomial, " ")
	a, _ := strconv.Atoi(splitString[0])
	operation := splitString[1]
	b, _ := strconv.Atoi(splitString[2])

	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	}

	return result
}