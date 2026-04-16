func twoSum(nums []int, target int) []int {
	front := 0
	rear  := len(nums) - 1

	var res []int

	for front < rear {
		sum := nums[front] + nums[rear]
		if sum == target {
			res = append(res, front+1)
			res = append(res, rear+1)
			break
		} else if sum > target {
			rear--
		} else {
			front++
		}
	}

	return res
}
