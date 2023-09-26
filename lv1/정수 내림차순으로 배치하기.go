func solution(n int64) int64 {

	s := strconv.FormatInt(n, 10)

	runeSlice := []rune(s)

	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] > runeSlice[j]
	})

	sortedStr := string(runeSlice)

	ret, _ := strconv.ParseInt(sortedStr, 10, 64)

	return ret
}