package main

import "fmt"

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
*/
// topN问题，可以用快排和堆排来解
func main() {
	nums := []int{1,1,1,1,1,1}
	k := 3
	fmt.Println(SelectTopK(nums, k, 0, len(nums) - 1))
}

func Partition(nums []int, low, high int) int {
	tmp := nums[low]
	for low < high {
		for low < high && nums[high] >= tmp {
			high --
		}
		nums[low], nums[high] = nums[high], nums[low]
		for low < high && nums[low] <= tmp {
			low ++
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return low
}

func SelectTopK(nums []int, K, low, high int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if low >= high {
		return nums[low]
	}
	privot := Partition(nums, low, high)
	if privot == len(nums) - K {
		return nums[privot]
	} else if privot > K {
		return SelectTopK(nums, K, low, privot - 1)
	} else {
		return SelectTopK(nums, K, privot + 1, high)
	}
}

