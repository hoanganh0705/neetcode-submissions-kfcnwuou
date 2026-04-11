func twoSum(nums []int, target int) []int {
	m := make(map[int] int)
	result := make([]int, 2)

	for i:=0; i<len(nums); i++ {
		diff := target - nums[i]
		value, ok := m[diff]
		if ok {
			result[0] = i
			result[1] = value
			return result
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
