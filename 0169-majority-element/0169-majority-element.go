func majorityElement(nums []int) int {
    mapData := make(map[int]int)
    maxKey := 0
    maxValue := 0
    for _, v := range nums {
        mapData[v]++
    }
    for k, v := range mapData {
        if v > maxValue {
            maxKey = k
            maxValue = v
        }
    }
    return maxKey
}