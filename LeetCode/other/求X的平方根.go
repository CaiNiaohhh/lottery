package main

import "fmt"

// 需要注意的地方是最后的取值应该是l-1, 因为循环的退出条件是l>r，也就是l是刚好越过mid的那个数
func mySqrt(x int) int {
	l, r, mid := 0, x, 0
	for l <= r {
		mid = l + (r - l) >> 1
		if mid * mid == x {
			return mid
		}
		if mid * mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

func main() {
	fmt.Println(mySqrt(2))
}