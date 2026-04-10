func combinationSum(nums []int, target int) [][]int {
    var res [][]int
	var path []int

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy (temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i< len(nums); i++ {
			if nums[i] > target {
				break
			}

			path = append(path, nums[i])

			backtrack(i, target-nums[i])

			path = path[:len(path) - 1]
		}
	}

	backtrack(0, target)
	return res
}
