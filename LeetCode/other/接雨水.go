package main

// 左边最大的和右边最大的取小的一方减去当前值


func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	var left, right []int
	Max, ans := -1, 0
	// 计算右边的最大值
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > Max {
			Max = height[i]
		}
		right = append(right, Max)
	}
	Reverse(right)
	// 计算左边的最大值
	Max = -1
	for i := 0; i < len(height); i++ {
		if height[i] > Max {
			Max = height[i]
		}
		left = append(left, Max)
	}
	for i := 1; i < len(height) - 1; i++ {
		tmp := left[i] - right[i]
		if tmp < 0 {
			ans += left[i] - height[i]
		} else {
			ans += right[i] - height[i]
		}
	}
	return ans
}

func Reverse(list []int) {
	for i, j := 0, len(list) - 1; i < j; i, j = i + 1, j - 1 {
		list[i], list[j] = list[j], list[i]
	}
}

// 双指针

func _trap(height []int) (ans int) {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height) - 1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left ++
		} else {
			ans += rightMax - height[right]
			right --
		}
	}
	return
}