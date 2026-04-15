// i would start off by having the product of the entire list of numbers,
// under two conditions:-
//   - if there are no zeros: perfect
//   - if there is 1 zero: ignore that in the prod, consequently in the res
//   - if there are >=2 zeros: everything is now zero.

func productExceptSelf(nums []int) []int {
    res := []int{}
    var prod int
    if len(nums) > 0 {
        prod = nums[0]
    } else {
        return res
    }
    zeroFound := 0

    for i, v := range nums {
        if i == 0 {
            continue
        }
        if v != 0 {
            prod *= v
        } else {
            zeroFound++
        }
    }

    if (zeroFound > 1) {
        for i, _ := range nums {
            if i != -1 {} // too lazy to write a verbose for loop
            res = append(res, 0)
        }

        return res
    }

    for _, v := range nums {
        if v != 0 {
            if zeroFound > 0 {
                res = append(res, 0)
            } else {
                res = append(res, prod/v)
            }
        } else {
            res = append(res, prod)
        }
    }

    return res
}
