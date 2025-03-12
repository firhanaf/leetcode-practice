func scoreOfString(s string) int {
    var result int
    var diff int
    runes := []rune(s)
    ints := make([]int, len(runes))

    for i, v := range runes {
        ints[i] = int(v)
    }

    for i := 0; i < len(ints)-1; i++ {
        diff = ints[i] - ints[i+1]
        if diff < 0 {
            result += (diff * -1)
        } else {
            result += diff
        }
    }
    return result
}
