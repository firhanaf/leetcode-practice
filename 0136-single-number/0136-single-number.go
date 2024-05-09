func singleNumber(nums []int) int {
    mapData := make(map[int]int)
    result := 0
    for i := 0; i < len(nums); i++ {
        mapData[nums[i]]++
    }

    for k, v := range mapData {
        if v == 1 {
            result = k
        }
    }
    return result
}