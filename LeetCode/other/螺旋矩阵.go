package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix) - 1, 0, len(matrix[0]) - 1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top ++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right --
		// 因为第一个和第二个循环互不影响，且对下面的判断都是正向的，所以可以在下面做一个判断即可
		if top > bottom || left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom --
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left ++
	}
	return res
}
