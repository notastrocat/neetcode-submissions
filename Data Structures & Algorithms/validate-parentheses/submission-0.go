func isValid(s string) bool {
    var stack []rune
    table := map[rune]rune{')':'(', '}':'{', ']':'['}

    for _, c := range s {
        if ( c == '(' || c == '{' || c == '[' ) {
            stack = append(stack, c)
        }
        if ( c == ')' || c == '}' || c == ']' ) {
            if len(stack) == 0 {
                return false
            } else if table[c] != stack[len(stack)-1] {
                return false
            }
            // logic to `pop`
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
