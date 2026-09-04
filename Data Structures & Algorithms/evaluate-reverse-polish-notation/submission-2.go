func evalRPN(tokens []string) int {
	operations := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {}}
	
	stack := make([]int, 0)
	var cache [2]int

	for _, elem := range tokens {
		if _, exist := operations[elem]; exist {
			cache[0] = stack[len(stack) - 2]
			cache[1] = stack[len(stack) - 1]
			switch elem {
			case "+":
				stack[len(stack) - 2] = cache[0] + cache[1]
			case "-":
				stack[len(stack) - 2] = cache[0] - cache[1]
			case "*":
				stack[len(stack) - 2] = cache[0] * cache[1]
			case "/":
				stack[len(stack) - 2] = cache[0] / cache[1]
			}
			stack = stack[:len(stack) - 1]
		} else {
			num := int(elem[len(elem) - 1] - '0')
			degree := 10
			for index := len(elem) - 2; index >= 0; index-- {
				if elem[index] == '-' {
					num *= -1
				} else {
					num += int(elem[index] - '0') * degree
					degree *= 10
				}
			}
			stack = append(stack, num)
		}
	}

	return stack[0]
}
