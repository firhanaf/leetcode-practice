func isAnagram(s string, t string) bool {
    mapData1 := make(map[rune]int)
    mapData2 := make(map[rune]int)

    for _, v := range s {
        mapData1[v]++
    }

    for _, v := range t {
        mapData2[v]++
    }
    if len(mapData1) != len(mapData2) {
        return false
    }

    for char, count := range mapData1 {
        if mapData2[char] != count {
            return false
        }
    }
    return true
}