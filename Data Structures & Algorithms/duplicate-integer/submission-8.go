func hasDuplicate(nums []int) bool {
    backet := make(map[int]struct{}, 6)

    for _, elem := range nums {
        if _, exist := backet[elem]; exist {
            return true
        }
        backet[elem] = struct{}{}
    }

    return false
}
