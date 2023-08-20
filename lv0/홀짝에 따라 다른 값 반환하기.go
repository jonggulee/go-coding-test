import "math"

func solution(n int) int {
	if n%2 == 1 {
		var oddSum int
		for i := 1; i <= n; i += 2 {
			oddSum += i
		}
		return oddSum
	} else {
		var squareSum int
		var squareRoot int
		for i := 2; i <= n; i += 2 {
			squareRoot = int(math.Pow(float64(i), 2))
			squareSum += squareRoot
		}
		return squareSum
	}
}