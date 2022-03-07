package main

import "fmt"

/*
两数之和：假设只有一种情况
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	Map := make(map[int]int)
	for index, num := range nums {
		if val, ok := Map[target - num]; ok {
			return []int{val, index}
		}
		Map[num] = index
	}
	return []int{}
}

func main() {
	nums := []int{}
	target := 6
	fmt.Println(twoSum(nums, target))
}