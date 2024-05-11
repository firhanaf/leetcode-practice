func finalValueAfterOperations(operations []string) int {
    counter := 0
    for _, v := range operations {
        if v == "--X" || v == "X--" {
            counter--
        }
        if v == "++X" || v == "X++" {
            counter++
        }
    }
    return counter
}