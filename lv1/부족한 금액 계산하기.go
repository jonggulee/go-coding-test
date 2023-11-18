func solution(price int, money int, count int) int64 {
	sum := 0

	for i := 1; i <= count; i++ {
		sum += (price * i)
	}

	if sum < money {
		return 0
	} else {
		return int64(sum - money)
	}
}