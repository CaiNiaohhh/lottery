package main
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 递归的解法会超时
func _climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n - 1) + climbStairs(n - 2)
}

// 尝试非递归的解法 斐波那契数列，初始值是0，0，1，循环n次
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
