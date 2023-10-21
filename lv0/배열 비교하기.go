func solution(arr1 []int, arr2 []int) int {
	sumArr1 := 0
	sumArr2 := 0

	if len(arr1) > len(arr2) {
		return 1
	} else if len(arr1) < len(arr2) {
		return -1
	} else {
		for i := 0; i < len(arr1); i++ {
			sumArr1 += arr1[i]
			sumArr2 += arr2[i]
		}

		if sumArr1 == sumArr2 {
			return 0
		} else if sumArr1 > sumArr2 {
			return 1
		} else {
			return -1
		}
	}
}