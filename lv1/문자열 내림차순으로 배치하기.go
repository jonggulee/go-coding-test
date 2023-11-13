import "sort"

func solution(s string) string {
	bytes := []byte(s)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})

	return string(bytes)
}