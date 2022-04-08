package other

// 回溯法

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	res, path, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) >= len(nums) {
			var tmp = make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for index := range nums {
			if used[index] {
				continue
			}
			path = append(path, nums[index])
			used[index] = true
			dfs()
			used[index] = false
			path = path[:len(path) - 1]
		}
	}
	dfs()
	return res
}
