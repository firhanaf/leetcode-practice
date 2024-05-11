func getConcatenation(nums []int) []int {
    var result []int
    result = nums
    for _, v := range nums {
        result = append(result, v)
    }
    return result
}