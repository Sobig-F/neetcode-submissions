func isPalindrome(s string) bool {
	left_index := 0
	right_index := len(s) - 1

	for range len(s) {
		if left_index >= right_index {
			break
		}

		for left_index < right_index &&
			(s[left_index] < '0' ||
			(s[left_index] > '9' && s[left_index] < 'A') ||
			(s[left_index] > 'Z' && s[left_index] < 'a') ||
			s[left_index] > 'z') {
			left_index++
		}
		for left_index < right_index &&
			(s[right_index] < '0' ||
			(s[right_index] > '9' && s[right_index] < 'A') ||
			(s[right_index] > 'Z' && s[right_index] < 'a') ||
			s[right_index] > 'z') {
			right_index--
		}

		a := s[left_index]
		b := s[right_index]

		if a >= 'A' && a <= 'Z'  {
			a += 32 
		}

		if b >= 'A' && b <= 'Z'  {
			b += 32
		}

		if a != b {
			return false
		}

		left_index++
		right_index--
	}

	return true
}
