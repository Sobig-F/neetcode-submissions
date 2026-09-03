func twoSum(numbers []int, target int) []int {
	for left, right := 0, len(numbers) - 1;; {
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
