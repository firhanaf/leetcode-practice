func differenceOfSums(n int, m int) int {
    var result int
    num1 := 0
    num2 := 0
    for i := 1; i <= n; i++ {
        if i % m != 0 {
            num1 += i
        }
        if i % m == 0 {
            num2 += i
        }
    }
    result = num1 - num2
    return result
}