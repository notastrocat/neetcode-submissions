// import "strconv"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder

    for _, str := range strs {
        sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteString("#")
        sb.WriteString(str)
    }

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    i := 0

    for i<len(encoded) {
        j := i
        for encoded[j] != '#' {
            j++
        }

        length, _ := strconv.Atoi(encoded[i:j])

        res = append(res, encoded[(j+1) : (length+j+1)])

        // how to increase i now? can't do ++
        i = j+1+length
    }

    return res
}
