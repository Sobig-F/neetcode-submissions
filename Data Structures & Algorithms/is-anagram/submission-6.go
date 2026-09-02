func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[rune]int)
	
	for i := range len(s) {
		set[rune(s[i])] += 1
		set[rune(t[i])] -= 1
	}

	for i := range set {
		if set[i] != 0 {
			return false
		}
	}

	return true
}
