func topKFrequent(nums []int, k int) []int {
    table := make(map[int]int)

    for _, num := range nums {
        table[num]++;
    }

    var res []int

    for n := range table {
        res = append(res, n)
    }

    sort.Slice(res, func(i, j int) bool {
        return table[res[i]] > table[res[j]]
    })

    return res[:k]
}
