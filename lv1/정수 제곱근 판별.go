import "math"

func solution(n int64) int64 {
	sqrt := math.Sqrt(float64(n))

	if sqrt == float64(int64(sqrt)) {
		return int64(sqrt+1) * int64(sqrt+1)
	}

	return -1
}