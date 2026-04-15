func topKFrequent(nums []int, k int) []int {
    table := make(map[int]int)
    maxFreq := 0
    for _, num := range nums {
        table[num]++
        if table[num] > maxFreq {
            maxFreq = table[num]
        }
    }

    // A slice of slices: index is frequency, value is a list of numbers
    buckets := make([][]int, maxFreq+1)
    for num, freq := range table {
        buckets[freq] = append(buckets[freq], num)
    }

    var res []int
    for i := maxFreq; i > 0 && len(res) < k; i-- {
        if len(buckets[i]) > 0 {
            res = append(res, buckets[i]...)
        }
    }

    // In case a bucket push us over 'k'
    return res[:k]
}