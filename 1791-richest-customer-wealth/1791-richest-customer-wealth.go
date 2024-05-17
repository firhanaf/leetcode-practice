func maximumWealth(accounts [][]int) int {
    result := 0
    sums := make([]int, len(accounts))
    for i, account := range accounts {
        accountSum := 0 
        for _, v := range account {
            accountSum += v
        }
        sums[i] = accountSum
    }
    for _, v := range sums{
        if v > result {
            result = v
        }
    }
    return result
}