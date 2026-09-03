func twoSum(numbers []int, target int) []int {
	result := make([]int, 0, 2)
	
	left := 0
	right := len(numbers) - 1

	for {
		if numbers[right] + numbers[left] == target {
			result = append(result, left + 1, right + 1)
			return result
		}

		if numbers[right] + numbers[left] > target {
			right--
		}

		if numbers[right] + numbers[left] < target {
			left++
		}
	}

	return result
}
