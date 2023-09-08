import "math"

func solution(brown int, yellow int) []int {
	total := brown + yellow

	for height := 1; height <= int(math.Sqrt(float64(total))); height++ {
		width := total / height

		if width >= height && (width-2)*(height-2) == yellow {
			return []int{width, height}
		}
	}

	return []int{}
}
