func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for {
		if numbers[right] + numbers[left] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[right] + numbers[left] > target {
			right--
		}
		if numbers[right] + numbers[left] < target {
			left++
		}

	}

	return nil
}
