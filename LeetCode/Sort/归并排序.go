package main

import (
	"fmt"
)

func mergeSort(arr []int) {
	tmp := make([]int, len(arr))
	Sort(arr, 0, len(arr) - 1, tmp)
}

func Sort(arr []int, left, right int, tmp []int) {
	if left < right {
		mid := left + (right - left) >> 1
		Sort(arr, 0, mid, tmp)
		Sort(arr, mid + 1, right, tmp)
		Merge(arr, left, mid, right, tmp)
	}
}

func Merge(arr []int, left, mid, right int, tmp []int) {
	i, j, cnt := left, mid + 1, 0
	for i <= mid || j <= right {
		if i <= mid && j <= right && arr[i] > arr[j] {	// 优先找出没有越界的数据进行比较
			tmp[cnt] = arr[j]
			cnt, j = cnt + 1, j + 1
		} else if i <= mid && j <= right && arr[i] <= arr[j] {	// 优先找出没有越界的数据进行比较
			tmp[cnt] = arr[i]
			cnt, i = cnt + 1, i + 1
		} else if i <= mid {	// 再找出越界的数据填充
			tmp[cnt] = arr[i]
			cnt, i = cnt + 1, i + 1
		} else if j <= right {	// 再找出越界的数据填充
			tmp[cnt] = arr[j]
			cnt, j = cnt + 1, j + 1
		}
	}
	cnt = 0
	for left <= right {
		arr[left] = tmp[cnt]
		cnt, left = cnt + 1, left + 1
	}
}

func main() {
	arr := []int{3,7,5,1,6,0}
	fmt.Println("before mergeSort：", arr)
	mergeSort(arr)
	fmt.Println("after mergeSort：", arr)
}