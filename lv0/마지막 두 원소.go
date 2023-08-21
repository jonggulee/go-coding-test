func solution(num_list []int) []int {
	numlistLen := len(num_list)
	if numlistLen < 2 {
		return num_list
	}

	lastElement := num_list[numlistLen-1]
	secondLastElement := num_list[numlistLen-2]

	if lastElement > secondLastElement {
		return append(num_list, lastElement-secondLastElement)
	} else {
		return append(num_list, lastElement*2)
	}
}