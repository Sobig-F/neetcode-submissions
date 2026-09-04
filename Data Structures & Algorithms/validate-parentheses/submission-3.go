func isValid(s string) bool {
    open_stack := make([]rune, 0, 5)
	close_bracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{'}

	for _, bracket := range s {
		if _, exist := close_bracket[bracket]; !exist {
			open_stack = append(open_stack, bracket)
		} else if len(open_stack) > 0 && close_bracket[bracket] == open_stack[len(open_stack) - 1] {
			open_stack = open_stack[:len(open_stack) - 1]
		} else {
			return false
		}
	}

	if len(open_stack) > 0 {
		return false
	}

	return true
}
