func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && result[i] == 0 && temperatures[stack[len(stack) - 1]] < temp {
			prevIdx := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result[prevIdx] = i - prevIdx
		}
		stack = append(stack, i)
	}

	return result
}
