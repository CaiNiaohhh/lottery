package main

import "fmt"

// 定义快排函数
func Partition(nums []int, low, high int) int {
	start, end := low, high
	tmp := nums[start]
	for start < end {
		for start < end && tmp <= nums[end] {
			end --
		}
		nums[start], nums[end] = nums[end], nums[start]
		for start < end && tmp >= nums[start] {
			start ++
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return start
}

func QuickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	privot := Partition(nums, low, high)
	fmt.Println(privot, low, high, nums)
	if low < privot {
		QuickSort(nums, low, privot - 1)
	}
	if privot < high {
		QuickSort(nums, privot + 1, high)
	}
}

func main() {
	nums := []int{3,4,10,7,9,8,5,6,2,1}
	QuickSort(nums, 0, len(nums) - 1)
	fmt.Println(nums)
}