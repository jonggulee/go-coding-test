import "sort"

func solution(rank []int, attendance []bool) int {
	rankMap := make(map[int]bool)
	rankIndex := make(map[int]int)
	top := []int{}
	for i, v := range rank {
		rankMap[v] = attendance[i]
		rankIndex[v] = i
	}

	for k, v := range rankMap {
		if v {
			top = append(top, k)
		}
	}
	sort.Ints(top)

	return (10000 * rankIndex[top[0]]) + (100 * rankIndex[top[1]]) + rankIndex[top[2]]
}