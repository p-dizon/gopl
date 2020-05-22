package duplicate

func removeAdjacentDuplicates(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			continue
		}

		copy(s[i:], s[i+1:])
		s = s[:len(s)-1]
	}
	return s
}
