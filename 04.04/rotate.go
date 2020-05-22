package rotate

func rotateLeft(s []int, amount int) []int {
	amount %= len(s)
	if amount == 0 {
		return s
	}
	s = append(s, s[:amount]...)
	return s[amount:]
}

func rotateRight(s []int, amount int) []int {
	amount %= len(s)
	amount = len(s) - amount
	return rotateLeft(s, amount)
}
