func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int)

	for _, num := range nums {
		set[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range set {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if buckets[i] != nil {
			result = append(result, buckets[i]...)
		}
	}

	return result
}
