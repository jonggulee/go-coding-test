import "strconv"

func solution(a string, b string) string {
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	a, b = reverse(a), reverse(b)
	lenA, lenB := len(a), len(b)
	carry := 0
	result := ""

	for i := 0; i < lenA || i < lenB || carry > 0; i++ {
		sum := carry
		if i < lenA {
			sum += int(a[i] - '0')
		}
		if i < lenB {
			sum += int(b[i] - '0')
		}
		carry = sum / 10
		result += strconv.Itoa(sum % 10)
	}

	return reverse(result)
}