func largestAltitude(gain []int) int {
    altitude := 0 
    result := 0
    sumAlt := []int{}
    sumAlt = append(sumAlt, altitude)
    for _, v := range gain {
        altitude += v
        sumAlt = append(sumAlt, altitude)
    }
    for _, v := range sumAlt {
        if v > result {
            result = v
        }
    }
    return result
}