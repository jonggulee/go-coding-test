import "sort"

func solution(num_list []int) []int {
	res := make([]int, 5)

	sort.Ints(num_list)

	for i := 0; i < 5; i++ {
		res[i] = num_list[i]
	}

	return res
}