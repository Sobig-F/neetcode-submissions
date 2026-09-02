func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := range nums[:len(nums) - 1] {
		for j := range len(nums) - i - 1 {
			if nums[i] + nums[i + 1 + j] == target {
				result[0] = i
				result[1] = i + 1 + j
				return result
			}
		}
	}
	return result
}
