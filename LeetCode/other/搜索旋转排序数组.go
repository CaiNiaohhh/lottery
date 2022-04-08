package other

/*
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
*/

func _search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 接下来是分区间，看哪个区间是有序的
		if nums[left] <= nums[mid] { // 如果是左区间有序
			if nums[left] <= target && target <= nums[mid] { // 如果target落在有序的区间
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 如果是右区间有序
			if nums[right] >= target && target >= nums[mid] { // 如果target落在有序的区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
