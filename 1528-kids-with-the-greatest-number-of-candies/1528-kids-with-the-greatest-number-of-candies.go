func kidsWithCandies(candies []int, extraCandies int) []bool {
    currentGreatCandies := 0
    afterExtraCandies := []int{}
    var result []bool
    for _, v := range candies {
        if v > currentGreatCandies {
            currentGreatCandies = v
        }
    }
    for _, v := range candies {
        afterExtraCandies = append(afterExtraCandies, v+extraCandies)
    }
    for _, v := range afterExtraCandies {
        if v >= currentGreatCandies {
            result = append(result, true)
        } else {
            result = append(result, false)
        }
    }
    return result
}