func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    counter := 0
    for _, v := range hours {
        if v >= target {
            counter++
        }
    }
    return counter
}