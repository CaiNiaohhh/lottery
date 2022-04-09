package main

// 使用双栈实现队列
// 使用切片模拟stack，具体是后进先出，始终操作最后一个元素

type MyQueue struct {
	inStack, outStack []int
}

func _Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	val := q.Peek()
	if val != -1 {
		q.outStack = q.outStack[:len(q.outStack) - 1]
	}
	return val
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack) - 1])
			q.inStack = q.inStack[:len(q.inStack) - 1]
		}
	}
	return q.outStack[len(q.outStack) - 1]
}

func (q *MyQueue) Empty() bool {
	if len(q.inStack) == 0 && len(q.outStack) == 0 {
		return true
	}
	return false
}



