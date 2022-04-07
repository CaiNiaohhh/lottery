package main

import "fmt"

/*
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
*/

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j] + 1)
				ans = Max(ans, dp[i])
			}

		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
}