func solution(number []int) int {
	count := 0

	for str := 0; str < len(number)-2; str++ {
		for mid := str + 1; mid < len(number)-1; mid++ {
			for end := mid + 1; end < len(number); end++ {
				if (number[str] + number[mid] + number[end]) == 0 {
					count += 1
				}
			}
		}
	}
	return count
}