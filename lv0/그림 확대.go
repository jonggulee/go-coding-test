func solution(picture []string, k int) []string {
	var res []string
	for _, row := range picture {
		expandedRow := ""
		for _, pixel := range row {
			for i := 0; i < k; i++ {
				expandedRow += string(pixel)
			}
		}

		for i := 0; i < k; i++ {
			res = append(res, expandedRow)
		}
	}
	return res
}