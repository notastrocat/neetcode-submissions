func groupAnagrams(strs []string) [][]string {
    table := make(map[[26]int][]string)  // we'll hold the key as the freq count
                                         // and the value can be the slice of strings

    // this approach will be an improvement on my last submission but only caters to
    // the English alphabets.
    for _, str := range strs {
        var freq [26]int
        for i, _ := range str {
            freq[str[i] - 'a']++
        }

        table[freq] = append(table[freq], str)
    }

    var results [][]string

    for _, v := range table {
        results = append(results, v)
    }

    return results
}
