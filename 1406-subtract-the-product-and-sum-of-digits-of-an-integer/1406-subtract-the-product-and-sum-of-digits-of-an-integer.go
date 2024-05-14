func subtractProductAndSum(n int) int {
    var splitInt []int
    for n > 0 {
        splitInt = append(splitInt, n%10)
        n /= 10
    }
    multiple := 1
    sum := 0
    for _, v := range splitInt {
        multiple *= v
    }
    for _, v := range splitInt {
        sum += v
    }
    return multiple - sum
}