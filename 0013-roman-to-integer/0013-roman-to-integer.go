func romanToInt(s string) int {
    integer := make(map[string]int)
    integer["I"] = 1
    integer["IV"] = 4
    integer["V"] = 5
    integer["IX"] = 9
    integer["X"] = 10
    integer["XL"] = 40
    integer["L"] = 50
    integer["XC"] = 90
    integer["C"] = 100
    integer["CD"] = 400
    integer["D"] = 500
    integer["CM"] = 900
    integer["M"] = 1000

    total := 0

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && integer[s[i:i+1]] < integer[s[i+1:i+2]] {
            total += integer[s[i:i+2]]
            i++
        } else {
            total += integer[s[i:i+1]]
        }
    }
    return total
}