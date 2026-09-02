func hasDuplicate(nums []int) bool {
    backet := make(map[int]int, len(nums))

    for _, elem := range nums {
        backet[elem] = elem
    }

    if len(backet) < len(nums) {
        return true
    }

    return false
}
