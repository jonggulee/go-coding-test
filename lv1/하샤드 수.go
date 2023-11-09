func solution(x int) bool {
	original := x
	sumOfDigits := 0

	for x > 0 {
		fmt.Println(x, sumOfDigits)
		sumOfDigits += x % 10
		x /= 10
	}

	return original%sumOfDigits == 0
}