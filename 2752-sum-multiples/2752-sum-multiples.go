func sumOfMultiples(n int) int {
    var result []int
    sum := 0
    for i := 1; i <= n; i++ {
        if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
            result = append(result, i)
        }
    }
    for _, v := range result {
        sum += v
    }
    return sum
}