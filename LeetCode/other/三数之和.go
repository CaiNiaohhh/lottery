package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？
请你找出所有和为 0 且不重复的三元组。注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]
*/

func MythreeSum(nums []int) [][]int{
	// 方法：排序后双指针
	length := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for index := 0; index < length; index++ {
		if index > 0 && nums[index] == nums[index - 1] {
			continue
		}
		target := -1 * nums[index]
		se, th := index + 1, length - 1
		for se <= th {
			source := nums[se] + nums[th]
			if source == target {
				res := []int{nums[index], nums[se], nums[th]}
				ans = append(ans, res)
				se++
				th--
			} else if source > target {
				th--
			} else {
				se++
			}
		}
	}
	return ans
}

func main() {
	var nums = []int{1,2,-3}
	fmt.Println(MythreeSum(nums))
}


