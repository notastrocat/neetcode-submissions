import (
	"slices"
	"cmp"
)

type compval struct {
	timestamp int
	val string
}

type TimeMap struct {
	table map[string][]compval
}

func Constructor() TimeMap {
	return TimeMap{
		table: make(map[string][]compval),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.table[key] = append(this.table[key], compval{
		timestamp: timestamp,
		val: value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list, ok := this.table[key]
    if !ok || len(list) == 0 {
        return ""
    }

    idx, found := slices.BinarySearchFunc(list, timestamp, func(c compval, targetTS int) int {
        return cmp.Compare(c.timestamp, targetTS)
    })

    if found {
        return list[idx].val
    }

    if idx > 0 {
        return list[idx-1].val
    }

    return ""
}