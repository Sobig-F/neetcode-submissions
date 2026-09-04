func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		} else {
			return 0
		}
	}

	i := len(nums) / 2 - 1

	if nums[i] == target {
		return i
	} else if nums[i] > target {
		result := search(nums[:i], target)
		return result
	} else if nums[i] < target {
		result := search(nums[i + 1:], target)
		if result == -1 {
			return -1
		}
		return i + 1 + result
	}

	return -1
}
