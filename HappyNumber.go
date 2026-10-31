package dsastringsarrays

func HappyNumber(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true
		sum := 0
		for n > 9 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}
	return true
}
