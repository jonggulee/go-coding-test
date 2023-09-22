import "math"

func solution(num_list []int) int {
	odd := 0
	even := 0
	for i := 0; i < len(num_list); i++ {
		if i%2 == 0 {
			even += num_list[i]
		} else {
			odd += num_list[i]
		}
	}

	if even > odd {
		return even
	} else {
		return odd
	}
}