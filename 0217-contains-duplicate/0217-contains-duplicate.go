func containsDuplicate(nums []int) bool {
    mapData := make(map[int]int)
    for _, v := range nums{
        mapData[v]++
    }
    for _, v := range mapData{
        if v >= 2 {
            return true
        }
    }
    return false
}