func solution(a int, d int, included []bool) int {
	var seq = a
	var r int
	for i := 0; i < len(included); i++ {
		if included[i] {
			r = r + seq
		}
		seq = seq + d
	}
	return r
}