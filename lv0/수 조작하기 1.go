func solution(n int, control string) int {
	for i := 0; i < len(control); i++ {
		switch control[i] {
		case 'w':
			n += 1
		case 's':
			n += -1
		case 'd':
			n += 10
		case 'a':
			n += -10
		}
	}
	return n
}