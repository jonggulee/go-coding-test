func solution(a int, b int, c int, d int) int {
	count := make(map[int]int)

	count[a]++
	count[b]++
	count[c]++
	count[d]++

	var uniqueNumbers []int
	for num := range count {
		uniqueNumbers = append(uniqueNumbers, num)
	}

	switch len(count) {
	case 1:
		return 1111 * a
	case 2:
		if count[uniqueNumbers[0]] == 3 || count[uniqueNumbers[1]] == 3 {
			major, minor := uniqueNumbers[0], uniqueNumbers[1]
			if count[uniqueNumbers[0]] == 1 {
				major, minor = minor, major
			}
			return int(math.Pow(float64(10*major+minor), 2))
		} else {
			return (uniqueNumbers[0] + uniqueNumbers[1]) * int(math.Abs(float64(uniqueNumbers[0]-uniqueNumbers[1])))
		}
	case 3:
		var single1, single2 int
		for _, num := range uniqueNumbers {
			if count[num] == 1 {
				if single1 == 0 {
					single1 = num
				} else {
					single2 = num
				}
			}
		}
		return single1 * single2
	case 4:
		sort.Ints(uniqueNumbers)
		return uniqueNumbers[0]
	}

	return 0
}