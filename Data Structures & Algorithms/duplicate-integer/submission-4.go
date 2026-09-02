func hasDuplicate(nums []int) bool {
    backet := make(map[int]interface{}, len(nums))

    for _, elem := range nums {
        backet[elem] = nil
    }

    if len(backet) < len(nums) {
        return true
    }

    return false
}
