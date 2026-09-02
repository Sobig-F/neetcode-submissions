func isAnagram(s string, t string) bool {
	set := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for i := range len(s) {
		if _, exist := set[rune(s[i])]; !exist {
			set[rune(s[i])] = 0 
		}
		set[rune(s[i])] += 1
		
		if _, exist := set[rune(t[i])]; !exist {
			set[rune(t[i])] = -1
		} else {
			set[rune(t[i])] -= 1
		}

	}

	// for _, elem := range s {
	// 	if _, exist := set[rune(elem)]; !exist {
	// 		set[rune(elem)] = 0 
	// 	}
	// 	set[rune(elem)] += 1 
	// }

	// for _, elem := range t {
	// 	if _, exist := set[rune(elem)]; !exist {
	// 		return false
	// 	}
	// 	set[rune(elem)] -= 1
	// }

	for i := range set {
		if set[i] != 0 {
			return false
		}
	}

	return true
}
