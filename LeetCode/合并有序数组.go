package main

/*
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, index := m - 1, n - 1, m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
	for i >= 0 {
		nums1[index] = nums1[i]
		index--
		i--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}