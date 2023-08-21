import "strconv"

func solution(num_list []int) int {
	var odd string
	var even string
	for i := 0; i < len(num_list); i++ {
		if num_list[i]%2 == 0 {
			even += strconv.Itoa(num_list[i])
		} else {
			odd += strconv.Itoa(num_list[i])
		}
	}
	oddInt, _ := strconv.Atoi(odd)
	evenInt, _ := strconv.Atoi(even)
	return oddInt + evenInt
}