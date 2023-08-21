import "math"

func solution(a int, b int, c int) int {
	var square2 = float64(2)
	var square3 = float64(3)
	if a == b && a == c && b == c {
		return (a + b + c) * (int(math.Pow(float64(a), square2)) + int(math.Pow(float64(b), square2)) + int(math.Pow(float64(c), square2))) * (int(math.Pow(float64(a), square3)) + int(math.Pow(float64(b), square3)) + int(math.Pow(float64(c), square3)))
	} else if a == b || a == c || b == c {
		return (a + b + c) * (int(math.Pow(float64(a), square2)) + int(math.Pow(float64(b), square2)) + int(math.Pow(float64(c), square2)))
	} else {
		return a + b + c
	}
}