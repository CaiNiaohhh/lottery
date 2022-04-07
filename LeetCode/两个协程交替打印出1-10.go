package main

import (
"fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	// 创建3个channel,A,B和Exit
	A := make(chan bool)
	B := make(chan bool)
	Exit := make(chan bool)  // 为了防止主进程退出

	go func() {
		// 如果A通道是true,我就执行
		for i := 1; i <= 9; i += 2 {
			if ok := <-A; ok {
				fmt.Println("A 输出", i, nums[i-1])
				B <- true
			}
		}
	}()

	go func() {
		defer func() { Exit <- true }() // 这个协程的活干完之后，向主goroutine发送信号
		// 如果B通道是true,我就执行
		for i := 2; i <= 10; i += 2 {
			if ok := <-B; ok {
				fmt.Println("B 输出", i, nums[i-1])
				if i != 10 { // r如果i等于10了，就不要再向A通道写数据了，否则将导致A通道死锁，至于为什么，坦白说我很疑惑
					A <- true
				}
				//A <- true
			}
		}
	}()

	A <- true // 启动条件
	<-Exit    // 结束条件
}
