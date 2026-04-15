func isPalindrome(s string) bool {
    var finalStr string
    for _, c := range s {
        if ( unicode.IsDigit(c) || unicode.IsLetter(c) ) {
            finalStr += string(unicode.ToLower(c))
        }
    }

    first := 0
    last := len(finalStr) - 1

    for first <= last {
        if finalStr[first] != finalStr[last] {
            return false
        }

        first++
        last--
    }

    return true
}
