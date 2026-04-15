import ("slices")

func groupAnagrams(strs []string) [][]string {
    table := make(map[string][]string)  // we'll hold the key as the sorted string and the
                                        // value can be the slice of strings

    for _, str := range strs {
        r := []rune(str)
        slices.Sort(r)

        table[string(r)] = append(table[string(r)], str)
    }

    var results [][]string

    for _, v := range table {
        results = append(results, v)
    }

    return results
}
