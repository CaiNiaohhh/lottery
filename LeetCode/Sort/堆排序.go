package main

import "fmt"

// len(nums) / 2 -> 第一个非叶子节点

func BuildMaxHeap(nums []int, length, index int) {
	tmp, left, right := index, index * 2 + 1, index * 2 + 2
	if left < length && nums[left] >= nums[tmp] {
		tmp = left
	}
	if right < length && nums[right] >= nums[tmp] {
		tmp = right
	}
	if tmp != index {
		nums[tmp], nums[index] = nums[index], nums[tmp]
		BuildMaxHeap(nums, length, tmp)
	}
}

func HeapSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	length := len(nums)
	// 从第一个非叶子节点开始创建最大堆
	cur := len(nums) / 2
	for i := cur; i >= 0; i-- {
		BuildMaxHeap(nums, length, i)
	}
	// 开始排序
	for i := length - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		BuildMaxHeap(nums, i, 0)
	}
}

func main() {
	nums := []int{4,3,2,1}
	HeapSort(nums)
	fmt.Println(nums)
}
