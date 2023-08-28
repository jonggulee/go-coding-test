import "strconv"

func solution(intStrs []string, k int, s int, l int) []int {
	r := []int{}
	for _, str := range intStrs {
		str = str[s : s+l]
		num, _ := strconv.Atoi(str)
		if num > k {
			r = append(r, num)
		}
	}
	return r
}